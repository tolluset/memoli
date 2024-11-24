package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type Memo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func main() {
	memos := loadFromFile()

	scanner := bufio.NewScanner(os.Stdin)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("Interrupted")
		saveToFile(memos)
		os.Exit(1)
	}()

	for {
		fmt.Print("Enter command: ")
		scanner.Scan()
		text := scanner.Text()

		if text == "exit" {
			break
		}

		if text == "add" {
			fmt.Print("Enter title: ")
			scanner.Scan()
			title := scanner.Text()
			fmt.Print("Enter content: ")
			scanner.Scan()
			content := scanner.Text()

			now := time.Now()
			formattedNow := now.Format("2006-01-02 15:04:05")

			newMemo := Memo{
				ID:        strconv.Itoa(len(memos) + 1),
				Title:     title,
				Content:   content,
				CreatedAt: formattedNow,
				UpdatedAt: formattedNow,
			}

			memos = AddMemo(memos, newMemo)
		}

		if text == "get" {
			memos = GetMemos(memos)

			for _, memo := range memos {
				fmt.Printf("%+v\n", memo)
			}
		}
	}

	saveToFile(memos)

	fmt.Print(memos)
	fmt.Printf("%+v\n", memos)
}

func saveToFile(memos []Memo) {
	file, err := os.Create("data/memo.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	for _, memo := range memos {
		fmt.Fprintf(file, "%s,%s,%s,%s,%s\n", memo.ID, memo.Title, memo.Content, memo.CreatedAt, memo.UpdatedAt)
	}
}

func loadFromFile() []Memo {
	file, err := os.Open("memo.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var memos []Memo
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		memo := Memo{
			ID:        fields[0],
			Title:     fields[1],
			Content:   fields[2],
			CreatedAt: fields[3],
			UpdatedAt: fields[4],
		}
		memos = append(memos, memo)
	}
	return memos
}

func AddMemo(memos []Memo, newMemo Memo) []Memo {
	memos = append(memos, newMemo)
	return memos
}

func GetMemos(memos []Memo) []Memo {
	return memos
}

func GetMemo(memos []Memo, id string) Memo {
	for _, memo := range memos {
		if memo.ID == id {
			return memo
		}
	}
	return Memo{}
}

func UpdateMemo(memos []Memo, updatedMemo Memo) []Memo {
	for i, memo := range memos {
		if memo.ID == updatedMemo.ID {
			memos[i] = updatedMemo
			break
		}
	}
	return memos
}

func DeleteMemo(memos []Memo, id string) []Memo {
	for i, memo := range memos {
		if memo.ID == id {
			memos = append(memos[:i], memos[i+1:]...)
			break
		}
	}
	return memos
}
