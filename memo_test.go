package main

import "testing"

func TestAddMemo(t *testing.T) {
	memos := []Memo{
		{
			ID:        "1",
			Title:     "title1",
			Content:   "content1",
			CreatedAt: "2021-01-01",
			UpdatedAt: "2021-01-01",
		},
		{
			ID:        "2",
			Title:     "title2",
			Content:   "content2",
			CreatedAt: "2021-01-02",
			UpdatedAt: "2021-01-02",
		},
	}
	newMemo := Memo{
		ID:        "3",
		Title:     "title3",
		Content:   "content3",
		CreatedAt: "2021-01-03",
		UpdatedAt: "2021-01-03",
	}
	got := AddMemo(memos, newMemo)
	if len(got) != 3 {
		t.Errorf("AddMemo() = %d; want 3", len(got))
	}
}

func TestGetMemos(t *testing.T) {
	memos := []Memo{
		{
			ID:        "1",
			Title:     "title1",
			Content:   "content1",
			CreatedAt: "2021-01-01",
			UpdatedAt: "2021-01-01",
		},
		{
			ID:        "2",
			Title:     "title2",
			Content:   "content2",
			CreatedAt: "2021-01-02",
			UpdatedAt: "2021-01-02",
		},
	}
	got := GetMemos(memos)
	if len(got) != 2 {
		t.Errorf("GetMemos() = %d; want 2", len(got))
	}
}

func TestGetMemo(t *testing.T) {
	memos := []Memo{
		{
			ID:        "1",
			Title:     "title1",
			Content:   "content1",
			CreatedAt: "2021-01-01",
			UpdatedAt: "2021-01-01",
		},
	}

	tests := []struct {
		ID   string
		want string
	}{
		{"1", "1"},
		{"2", ""},
	}

	for _, test := range tests {
		got := GetMemo(memos, test.ID)
		if got.ID != test.want {
			t.Errorf("GetMemo() = %s; want %s", got.ID, test.want)
		}
	}
}

func TestUpdateMemo(t *testing.T) {
	memos := []Memo{
		{
			ID:        "1",
			Title:     "title1",
			Content:   "content1",
			CreatedAt: "2021-01-01",
			UpdatedAt: "2021-01-01",
		},
		{
			ID:        "2",
			Title:     "title2",
			Content:   "content2",
			CreatedAt: "2021-01-02",
			UpdatedAt: "2021-01-02",
		},
	}
	updatedMemo := Memo{
		ID:        "1",
		Title:     "title1",
		Content:   "content1",
		CreatedAt: "2021-01-01",
		UpdatedAt: "2021-01-03",
	}
	got := UpdateMemo(memos, updatedMemo)
	if got[0].UpdatedAt != "2021-01-03" {
		t.Errorf("UpdateMemo() = %s; want 2021-01-03", got[0].UpdatedAt)
	}
}

func TestDeleteMemo(t *testing.T) {
	memos := []Memo{
		{
			ID:        "1",
			Title:     "title1",
			Content:   "content1",
			CreatedAt: "2021-01-01",
			UpdatedAt: "2021-01-01",
		},
		{
			ID:        "2",
			Title:     "title2",
			Content:   "content2",
			CreatedAt: "2021-01-02",
			UpdatedAt: "2021-01-02",
		},
	}
	got := DeleteMemo(memos, "1")
	if len(got) != 1 {
		t.Errorf("DeleteMemo() = %d; want 1", len(got))
	}
}
