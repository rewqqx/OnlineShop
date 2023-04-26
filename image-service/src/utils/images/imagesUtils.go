package images

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var pathToImages = "./resources/images"

func getImagePathByID(directory string, id int) (string, error) {
	files, err := filepath.Glob(directory + "/*")
	if err != nil {
		return "", err
	}

	for _, file := range files {
		filename := filepath.Base(file)
		if idString := fmt.Sprintf("%d", id); filename[:len(idString)] == idString {
			return file, nil
		}
	}

	return "", fmt.Errorf("image with ID %d not found", id)
}

func GetImageBytes(id int) ([]byte, error) {

	imagePath, err := getImagePathByID(pathToImages, id)
	if err != nil {
		return nil, err
	}

	imageBytes, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}

	return imageBytes, nil
}
