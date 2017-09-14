package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"vvoid.pw/archivebuffer"
	"vvoid.pw/randomizer"
)

const apiVersion = "v1"

var (
	apiRootPath = formatAPIPath("")
	apiPushPath = formatAPIPath("push")
)

var (
	serverPort     *int
	apiStoragePath *string
	apiSecret      *string
)

const (
	rootWelcomeMsg = "Welcome to the Pimba API."

	methodNotAllowed    = `{"error":"Method not allowed."}`
	badRequest          = `{"error":"Bad request."}`
	internalServerError = `{"error":"Internal Server Error. Please, try again later."}`
	unauthorized        = `{"error":"Token is invalid!"}`
)

func Serve(port int, storagePath, secret string) {
	serverPort = &port
	apiStoragePath = &storagePath
	apiSecret = &secret

	http.Handle("/", http.FileServer(http.Dir(*apiStoragePath)))

	http.HandleFunc(apiRootPath, rootHandler)
	http.HandleFunc(apiPushPath, pushHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *serverPort), nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, rootWelcomeMsg)
	return
}

func pushHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		w.Header().Set("Content-Type", "application/json")

		file, _, err := r.FormFile("file")
		if err != nil {
			log.Println("pushHandler:", err)
			http.Error(w, badRequest, http.StatusBadRequest)
			return
		}
		defer file.Close()

		ungzippedFile, err := archivebuffer.UngzipToBuffer(file)
		if err != nil {
			log.Println("pushHandler:", err)
			http.Error(w, internalServerError, http.StatusInternalServerError)
			return
		}

		var untarID string
		if r.FormValue("name") != "" {
			untarID = r.FormValue("name")
		} else {
			untarID = randomizer.GenerateRandomString(10)
		}
		untarPath := *apiStoragePath + "/" + untarID

		dirChecker, err := isDirExist(untarPath)
		if err != nil {
			http.Error(w, internalServerError, http.StatusInternalServerError)
			return
		}

		token := r.FormValue("token")
		if dirChecker && token == "" {
			http.Error(w, unauthorized, http.StatusUnauthorized)
			return
		} else if dirChecker {
			claims, err := ParseToken(token, *apiSecret)
			if err != nil {
				http.Error(w, unauthorized, http.StatusUnauthorized)
				return
			}
			if claims["bucket"] == untarID {
				err = archivebuffer.UntarToFile(ungzippedFile, untarPath)
				if err != nil {
					log.Println("pushHandler:", err)
					http.Error(w, internalServerError, http.StatusInternalServerError)
					return
				}
			} else {
				http.Error(w, unauthorized, http.StatusUnauthorized)
				return
			}
		} else {
			err = os.MkdirAll(untarPath, 0766)
			if err != nil {
				log.Println("untar:", err)
				http.Error(w, internalServerError, http.StatusInternalServerError)
				return
			}
			token, err = NewToken(untarID, *apiSecret)
			if err != nil {
				log.Println("pushHandler:", err)
				http.Error(w, internalServerError, http.StatusInternalServerError)
				return
			}

			err = archivebuffer.UntarToFile(ungzippedFile, untarPath)
			if err != nil {
				log.Println("pushHandler:", err)
				http.Error(w, internalServerError, http.StatusInternalServerError)
				return
			}
		}

		jsonResponse := fmt.Sprintf(`{ "push-url": "%v", "token": "%v" }`,
			r.Host+"/"+untarID, token)
		fmt.Fprintf(w, jsonResponse)
		return
	} else {
		http.Error(w, methodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
}

func formatAPIPath(apiCall string) string {
	return "/" + apiVersion + "/" + apiCall
}

func isDirExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
