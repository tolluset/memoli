package main

import "testing"

func TestAddMemo(t *testing.T) {
	memos := []Memo{
		{
			id:         "1",
			title:      "title1",
			content:    "content1",
			created_at: "2021-01-01",
			updated_at: "2021-01-01",
		},
		{
			id:         "2",
			title:      "title2",
			content:    "content2",
			created_at: "2021-01-02",
			updated_at: "2021-01-02",
		},
	}
	newMemo := Memo{
		id:         "3",
		title:      "title3",
		content:    "content3",
		created_at: "2021-01-03",
		updated_at: "2021-01-03",
	}
	got := AddMemo(memos, newMemo)
	if len(got) != 3 {
		t.Errorf("AddMemo() = %d; want 3", len(got))
	}
}

func TestGetMemos(t *testing.T) {
	memos := []Memo{
		{
			id:         "1",
			title:      "title1",
			content:    "content1",
			created_at: "2021-01-01",
			updated_at: "2021-01-01",
		},
		{
			id:         "2",
			title:      "title2",
			content:    "content2",
			created_at: "2021-01-02",
			updated_at: "2021-01-02",
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
			id:         "1",
			title:      "title1",
			content:    "content1",
			created_at: "2021-01-01",
			updated_at: "2021-01-01",
		},
	}

	tests := []struct {
		id   string
		want string
	}{
		{"1", "1"},
		{"2", ""},
	}

	for _, test := range tests {
		got := GetMemo(memos, test.id)
		if got.id != test.want {
			t.Errorf("GetMemo() = %s; want %s", got.id, test.want)
		}
	}
}

func TestUpdateMemo(t *testing.T) {
	memos := []Memo{
		{
			id:         "1",
			title:      "title1",
			content:    "content1",
			created_at: "2021-01-01",
			updated_at: "2021-01-01",
		},
		{
			id:         "2",
			title:      "title2",
			content:    "content2",
			created_at: "2021-01-02",
			updated_at: "2021-01-02",
		},
	}
	updatedMemo := Memo{
		id:         "1",
		title:      "title1",
		content:    "content1",
		created_at: "2021-01-01",
		updated_at: "2021-01-03",
	}
	got := UpdateMemo(memos, updatedMemo)
	if got[0].updated_at != "2021-01-03" {
		t.Errorf("UpdateMemo() = %s; want 2021-01-03", got[0].updated_at)
	}
}

func TestDeleteMemo(t *testing.T) {
	memos := []Memo{
		{
			id:         "1",
			title:      "title1",
			content:    "content1",
			created_at: "2021-01-01",
			updated_at: "2021-01-01",
		},
		{
			id:         "2",
			title:      "title2",
			content:    "content2",
			created_at: "2021-01-02",
			updated_at: "2021-01-02",
		},
	}
	got := DeleteMemo(memos, "1")
	if len(got) != 1 {
		t.Errorf("DeleteMemo() = %d; want 1", len(got))
	}
}
