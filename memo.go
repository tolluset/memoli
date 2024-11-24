package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	memos := []Memo{}

	scanner := bufio.NewScanner(os.Stdin)

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

	fmt.Print(memos)
	fmt.Printf("%+v\n", memos)
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
