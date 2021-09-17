package files

import (
	"log"
	"os"
	"path"
	"strings"
)

//-- Author : Ponsard Nils
//-- Last update : 10/08/2021

// Fonction : parse le chemin en remplaçant le ~ par le dossier home
func ParsePath(p string) string {

	//-- Compatibilité avec Windows

	parts := strings.Split(p, "/")
	joinedPath := path.Join(parts...)

	//-- on cherche le dossier de l’utilisateur

	homedir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("Could not find the home dir : ", err)
	}

	//-- On remplace le tilde par le dossier utilisateur

	return path.Clean(strings.ReplaceAll(joinedPath, "~", homedir))
}

//-- Author : Ponsard Nils
//-- Last update : 10/08/2021

// Fonction : s’assure que le dossier existe, le crée sinon
func EnsureFolder(p string) {
	base := path.Dir(p)
	os.MkdirAll(base, 0755)
}

//-- Author : Ponsard Nils
//-- Last update : 10/08/2021

// Fonction : vérifie si le fichier existe
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
