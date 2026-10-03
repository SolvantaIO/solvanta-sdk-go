package solvanta_test

import (
	"context"
	"fmt"
	"log"

	"github.com/SolvantaIO/solvanta-sdk-go"
)

func ExampleClient_Solve_recaptchaV3() {
	client := solvanta.New("sk-live-xxxxxxxx")

	resp, err := client.Solve(context.Background(), solvanta.ReCaptchaV3Task, map[string]interface{}{
		"url":     "https://example.com/login",
		"sitekey": "6LcXXXXXXXXXXXXXXXXXXXXXX",
		"action":  "login",
		"proxy":   "http://user:pass@proxy.example.com:8080",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status:", resp.Status)
	fmt.Println("Token:", string(resp.Result))
}

func ExampleClient_Solve_cloudflare() {
	client := solvanta.New("sk-live-xxxxxxxx")

	resp, err := client.Solve(context.Background(), solvanta.CloudflareTask, map[string]interface{}{
		"url":   "https://target-site.com",
		"proxy": "http://user:pass@proxy.example.com:8080",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status:", resp.Status)
	fmt.Println("Credits used:", resp.CreditsUsed)
	fmt.Println("Elapsed:", resp.ElapsedMs, "ms")
}

func ExampleClient_Solve_arkose() {
	client := solvanta.New("sk-live-xxxxxxxx")

	resp, err := client.Solve(context.Background(), solvanta.ArkoseTask, map[string]interface{}{
		"url":     "https://example.com",
		"sitekey": "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
		"proxy":   "http://user:pass@proxy.example.com:8080",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status:", resp.Status)
}

func ExampleClient_Solve_metadata1() {
	client := solvanta.New("sk-live-xxxxxxxx")

	resp, err := client.Solve(context.Background(), solvanta.Metadata1Task, map[string]interface{}{
		"email":    "user@example.com",
		"password": "hunter2",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status:", resp.Status)
}

func ExampleClient_SolveWithKey() {
	client := solvanta.New("sk-live-xxxxxxxx")

	resp, err := client.SolveWithKey(
		context.Background(),
		solvanta.ReCaptchaV3Task,
		map[string]interface{}{
			"url":     "https://example.com/login",
			"sitekey": "6LcXXXXXXXXXXXXXXXXXXXXXX",
			"action":  "login",
		},
		"order-12345-captcha",
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Solve ID:", resp.ID)
}

func ExampleClient_GetSolve() {
	client := solvanta.New("sk-live-xxxxxxxx")

	resp, err := client.GetSolve(context.Background(), "solve-uuid-here")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status:", resp.Status)
}

func ExampleClient_Services() {
	client := solvanta.New("sk-live-xxxxxxxx")

	resp, err := client.Services(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, svc := range resp.Services {
		fmt.Printf("%s — %s (%d credits)\n", svc.Task, svc.Description, svc.Credits)
	}
}

func ExampleClient_Dashboard() {
	client := solvanta.New("sk-live-xxxxxxxx")

	resp, err := client.Dashboard(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Credits: %d | Used today: %d | Solves today: %d\n",
		resp.Credits, resp.UsedToday, resp.SolvesToday)
}

func ExampleClient_errorHandling() {
	client := solvanta.New("sk-live-bad-key")

	_, err := client.Solve(context.Background(), solvanta.CloudflareTask, map[string]interface{}{
		"url": "https://example.com",
	})
	if err != nil {
		if apiErr, ok := err.(*solvanta.Error); ok {
			fmt.Println("API error:", apiErr.Code, "-", apiErr.Message)
			fmt.Println("HTTP status:", apiErr.StatusCode)
		} else {
			fmt.Println("Network error:", err)
		}
	}
}
