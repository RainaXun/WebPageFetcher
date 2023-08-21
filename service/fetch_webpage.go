package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func fetch(wg *sync.WaitGroup, path string, strategy string, URL string) error {
	defer wg.Done()

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	switch strategy {
	case fetchStrategySaveContent:
		parsedURL, _ := url.Parse(URL)
		file, err := os.Create(filepath.Join(path, parsedURL.Host+".html"))
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		defer file.Close()
		_, err = file.Write(body)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
	case fetchStrategyGetMetadata:
		dynamicMetadata, err := dynamicAnalysis(URL)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		staticMetadata, err := staticAnalysis(URL, resp)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		fmt.Println()
		PrintMetadata(Metadata{
			Site:      URL,
			NumLinks:  dynamicMetadata.NumLinks + staticMetadata.NumLinks,
			NumImages: dynamicMetadata.NumImages + staticMetadata.NumImages,
			LastFetch: resp.Header.Get("Date"),
		})
	default:
		return errors.New("invalid strategy")
	}
	return nil
}

func dynamicAnalysis(url string) (Metadata, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	if err := chromedp.Run(ctx, chromedp.Navigate(url)); err != nil {
		return Metadata{}, err
	}
	var linkCount, imageCount int
	if err := chromedp.Run(ctx,
		chromedp.EvaluateAsDevTools(`document.querySelectorAll('a').length`, &linkCount),
		chromedp.EvaluateAsDevTools(`document.querySelectorAll('img').length`, &imageCount),
	); err != nil {
		return Metadata{}, err
	}

	return Metadata{
		Site:      url,
		NumLinks:  linkCount,
		NumImages: imageCount,
		LastFetch: "",
	}, nil
}

func staticAnalysis(url string, resp *http.Response) (Metadata, error) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return Metadata{}, err
	}

	// Find and count links
	linkCount := doc.Find("a").Length()

	// Find and count images
	imageCount := doc.Find("img").Length()

	// Print the results
	return Metadata{
		Site:      url,
		NumLinks:  linkCount,
		NumImages: imageCount,
		LastFetch: resp.Header.Get("Date"),
	}, nil
}
