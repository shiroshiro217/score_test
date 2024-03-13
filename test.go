package main

import "fmt"

type Student_Data struct {
	name string
}

func (s Student_Data) calavg(scores []float64) (avg_score float64) {
	sum := 0.0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	avg_score = sum / float64(len(scores))
	return
}

func (s Student_Data) judge(score float64) {
	if score >= 60 {
		fmt.Println("考試結果及格了!")
	} else if 0 <= score && score < 60 {
		fmt.Println("考試結果不及格!")
	} else {
		fmt.Println("分數格式可能輸入錯誤了!")
	}
}

func main() {
	var s01 Student_Data

	subject := []string{"國文", "英文", "數學", "社會", "自然"}
	subject_score := []float64{0, 0, 0, 0, 0}

	fmt.Printf("請輸入你的名字:")
	fmt.Scanln(&s01.name)

	for i := 0; i < len(subject); i++ {
		fmt.Printf("請輸入%s的成績:", subject[i])
		fmt.Scanln(&subject_score[i])
	}

	result := s01.calavg(subject_score)
	fmt.Printf("%s同學,你的平均分數是:%f分\n", s01.name, result)
	s01.judge(result)
}
