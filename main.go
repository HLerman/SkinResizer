package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Println("Indiquer un chemin valide vers la skin à modifier")
		os.Exit(1)
	}

	path := args[0]

	if _, err := os.Stat(path); err != nil {
		log.Println("Chemin de fichier invalide")
		os.Exit(1)
	}

	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	imageName := filepath.Base(path)

	if img.Bounds().Dx() != 216 || img.Bounds().Dy() != 128 {
		log.Printf("L'image %s doit être au format 216x128", imageName)
		os.Exit(0)
	}

	/*palettedImg, ok := img.(*image.Paletted)
	if !ok {
		log.Printf("L'image %s ne contient pas de palette de couleurs", imageName)
		os.Exit(0)
	}*/

	// Afficher la palette de couleurs
	/*fmt.Println("Palette de couleurs :")
	c := palettedImg.Palette[0]
	fmt.Println(c.RGBA())*/
	/*for _, c := range palettedImg.Palette {
		//fmt.Printf(" - #%02x%02x%02x\n", c.R, c.G, c.B)
		fmt.Println(c.RGBA())
	}*/

	x := 0
	y := 0

	s := 4

	rgba := image.NewRGBA(image.Rect(0, 0, 288, 128))

	// On divise la skin en frame de 24 par 32
	for i := 0; i < 36; i++ {
		cropSize := image.Rect(0, 0, 24, 32)
		cropSize = cropSize.Add(image.Point{x, y})
		croppedImage := img.(SubImager).SubImage(cropSize)

		/*err := createFrame(croppedImage, fmt.Sprintf("%d.%s", i, imageName))
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}*/

		// On positionne la frame extraite dans la nouvelle image
		draw.Draw(rgba, image.Rect(s, y, s+24, y+32), croppedImage, image.Point{x, y}, draw.Src)

		y = y + 32

		if y == 128 {
			// S ici vaut 32 car on décale la frame non pas de 24px mais de 32 pour la positionner correctement dans la nouvelle image
			s = s + 32
			x = x + 24
			y = 0
		}
	}

	out, err := os.Create(imageName)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if err := png.Encode(out, rgba); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer out.Close()

	log.Printf("L'image %s a bien été générée", imageName)
}

/*func createFrame(croppedImage image.Image, nom string) error {
	croppedImageFile, err := os.Create("output/" + nom)
	if err != nil {
		return err
	}

	defer croppedImageFile.Close()
	if err := png.Encode(croppedImageFile, croppedImage); err != nil {
		return err
	}

	return nil
}*/
