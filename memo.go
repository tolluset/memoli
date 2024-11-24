package main

type Memo struct {
	id         string
	title      string
	content    string
	created_at string
	updated_at string
}

func main() {
	memos := []Memo{}

	print(memos)
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
		if memo.id == id {
			return memo
		}
	}
	return Memo{}
}

func UpdateMemo(memos []Memo, updatedMemo Memo) []Memo {
	for i, memo := range memos {
		if memo.id == updatedMemo.id {
			memos[i] = updatedMemo
			break
		}
	}
	return memos
}

func DeleteMemo(memos []Memo, id string) []Memo {
	for i, memo := range memos {
		if memo.id == id {
			memos = append(memos[:i], memos[i+1:]...)
			break
		}
	}
	return memos
}
