func CallAPI(req *http.Request) (sb interface{}, err error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return
	}
	err = json.Unmarshal(body, &sb)
	if err != nil {
		return
	}

	return
}
