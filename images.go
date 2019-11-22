package main

type Image struct {
	Image string
	URL   string
}

func getImages() map[string]*Image {
	images := map[string]*Image{
		"awkward": {
			Image: "https://media.giphy.com/media/kaq6GnxDlJaBq/giphy-downsized.gif",
			URL:   "https://giphy.com/gifs/chloe-concerned-kaq6GnxDlJaBq",
		},
		"calm down": {
			Image: "https://media.giphy.com/media/26uf7I0OKqyIpUO5O/giphy-downsized.gif",
			URL:   "https://giphy.com/gifs/viceprincipals-hbo-vice-principals-26uf7I0OKqyIpUO5O",
		},
		"disappointed": {
			Image: "https://media.giphy.com/media/3oAt21Fnr4i54uK8vK/giphy-downsized.gif",
			URL:   "https://giphy.com/gifs/hercules-disappointed-let-down-3oAt21Fnr4i54uK8vK/media",
		},
		"do it": {
			Image: "https://media.giphy.com/media/wi8Ez1mwRcKGI/giphy-downsized.gif",
			URL:   "https://giphy.com/gifs/elvira-georgia-dex-wi8Ez1mwRcKGI",
		},
		"impressive": {
			Image: "https://media.giphy.com/media/bJwba7vDxsBws/giphy-downsized.gif",
			URL:   "https://giphy.com/gifs/reaction-impressive-bJwba7vDxsBws",
		},
		"noice": {
			Image: "https://media.giphy.com/media/l3q2slcE8854Yqqg8/giphy-downsized.gif",
			URL:   "https://giphy.com/gifs/snl-saturday-night-live-season-42-l3q2slcE8854Yqqg8/",
		},
		"true dat": {
			Image: "https://media.giphy.com/media/udPdpF18yG0uI/giphy-downsized.gif",
			URL:   "https://giphy.com/gifs/true-dat-udPdpF18yG0uI",
		},
		"tubular": {
			Image: "https://media.giphy.com/media/3o7TKy1qgGdbbMalcQ/giphy-downsized.gif",
			URL:   "https://giphy.com/gifs/studiosoriginals-gloria-domitille-collardey-business-woman-3o7TKy1qgGdbbMalcQ",
		},
		"woop woop": {
			Image: "https://media.giphy.com/media/J0a9SREMHkBAA/giphy-downsized.gif",
			URL:   "https://giphy.com/gifs/J0a9SREMHkBAA",
		},
	}

	return images
}

func getImageKeys() []string {
	images := getImages()
	imageKeys := make([]string, 0, len(images))

	for k := range images {
		imageKeys = append(imageKeys, "• "+k)
	}

	return imageKeys
}

func getImage(input string) *Image {
	images := getImages()
	image, ok := images[input]

	if ok {
		return image
	}

	return nil
}
