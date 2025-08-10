package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Не найден файл .env")
	}
}

const TOKEN_URL = "https://web.vk.me/?act=web_token&app_id=8202606&v="
const VK_GROUP_URL = "https://api.vk.com/method/messages.getConversationMembers?&extended=1&fields=educational_profile&peer_id="

func main() {
	var conf = NewConfig()
	jar, _ := cookiejar.New(nil)
	var cookies []*http.Cookie
	cookie := &http.Cookie{
		Name:  "remixdsid",
		Value: conf.Sferum.REMIXDSID,
	}
	u, _ := url.Parse(TOKEN_URL)
	cookies = append(cookies, cookie)
	jar.SetCookies(u, cookies)

	client := &http.Client{
		Jar: jar,
	}
	var tokens Tokens
	err := fetchParseJson(client, TOKEN_URL+conf.Sferum.V, &tokens)
	if err != nil {
		fmt.Printf("Ошибка получения токена: %v\n", err)
		os.Exit(1)
	}

	var token VkToken = findSferumToken(tokens)
	if token.Token == "" {
		fmt.Println("Токен сферум не найден")
		os.Exit(1)
	}
	var profilesResponse VkProfileResponse
	err = fetchParseJson(client, VK_GROUP_URL+conf.Sferum.CHAT+"&access_token="+token.Token+"&v="+conf.Sferum.VVK, &profilesResponse)
	if err != nil {
		fmt.Printf("Ошибка получения данных: %v\n", err)
		os.Exit(1)
	}
	var profiles = profilesResponse.Response.Profiles

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"id", "ФИО", "Учебные заведения", "Роли"})
	tf := table.NewWriter()
	tf.AppendHeader(table.Row{"id", "ФИО", "Учебные заведения", "Роли"})
	for _, profile := range profiles {
		var organizations = ""
		for _, org := range profile.Eduprofile.Eduroles.Organizations {
			org.Name = splitString(org.Name, 30)
			organizations += org.Name + "\n"
		}
		organizations = strings.TrimSuffix(organizations, "\n")
		var rowData = []any{profile.Id, profile.Surname + " " + profile.Name, organizations, splitString(profile.Eduprofile.Eduroles.Mainroles, 25)}
		var rowDataFile = []any{profile.Id, profile.Surname + " " + profile.Name, strings.ReplaceAll(organizations, "\n", ""), strings.ReplaceAll(profile.Eduprofile.Eduroles.Mainroles, "\"", "'")}

		tw.AppendRow(rowData)
		tf.AppendRow(rowDataFile)
		tw.SetAllowedRowLength(150)
		tw.AppendSeparator()
	}
	fmt.Println(tw.Render())
	file, err := os.Create("output.csv")

	if err != nil {
		fmt.Println("невозможно создать файл:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(strings.ReplaceAll(tf.RenderCSV(), "\\", ""))
}
