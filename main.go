// This program will read the records from the CSV file. The CSV filename is passed to the program
// as a argument and will compute the following stats by gender:
// 	1. total number of records
// 	2. min salary
// 	3. max salary
// 	4. average salary

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	fname  string
	lname  string
	ssn    string
	gender string
	age    int
	salary float64
}

type People []Person

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func newPerson(f, l, ssn, g string, a int, s float64) Person {
	return Person{
		fname:  f,
		lname:  l,
		ssn:    ssn,
		gender: g,
		age:    a,
		salary: s,
	}
}

func getMales(d []Person) []Person {
	maleList := []Person{}
	for _, i := range d {
		if i.gender == "M" {
			maleList = append(maleList, i)
		}
	}
	return maleList
}

func getFemales(d []Person) []Person {
	femaleList := []Person{}
	for _, i := range d {
		if i.gender == "F" {
			femaleList = append(femaleList, i)
		}
	}
	return femaleList
}

func lowestFSalary(f []Person) float64 {
	var lowSalary float64 = 999999999
	for _, i := range f {
		if i.salary < lowSalary {
			lowSalary = i.salary
		}
	}
	return lowSalary
}

func MaxFSalary(f []Person) (float64, Person) {
	var maxSalary float64
	for _, i := range f {
		if i.salary > maxSalary {
			maxSalary = i.salary
		}
	}
	var maxPerson Person
	for _, i := range f {
		if i.salary == maxSalary {
			maxPerson = i
		}
	}
	return maxSalary, maxPerson
}

func averageFSalary(f []Person) float64 {
	var sum float64
	numSalaries := 0
	for _, i := range f {
		sum += i.salary
		numSalaries++
	}
	return sum / float64(numSalaries)
}

func lowestMSalary(m []Person) float64 {
	var lowSalary float64 = 999999999
	for _, i := range m {
		if i.salary < lowSalary {
			lowSalary = i.salary
		}
	}
	return lowSalary
}

func MaxMSalary(m []Person) (float64, Person) {
	var maxSalary float64
	for _, i := range m {
		if i.salary > maxSalary {
			maxSalary = i.salary
		}
	}
	var maxPerson Person
	for _, i := range m {
		if i.salary == maxSalary {
			maxPerson = i
		}
	}
	return maxSalary, maxPerson
}

func averageMSalary(m []Person) float64 {
	var sum float64
	numSalaries := 0
	for _, i := range m {
		sum += i.salary
		numSalaries++
	}
	return sum / float64(numSalaries)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Missing csv file: %v <csv_filename>\n", os.Args[0])
	}

	var db People

	csvRaw, err := ioutil.ReadFile(os.Args[1])
	checkError(err)

	csvData := string(csvRaw)
	csvData = strings.TrimSpace(csvData)
	csvData = strings.Replace(csvData, "\r\n", "\n", -1)
	csvDataLines := strings.Split(csvData, "\n")

	for _, i := range csvDataLines {
		person := strings.Split(i, ",")
		pfname := person[0]
		plname := person[1]
		pssn := person[2]
		pgender := person[3]
		page, _ := strconv.Atoi(person[4])
		psalary, _ := strconv.ParseFloat(person[5], 64)

		db = append(db, newPerson(pfname, plname, pssn, pgender, page, psalary))
	}

	// Get Slices for M/F
	Females := getFemales(db)
	Males := getMales(db)

	// Print Female Stats
	fmt.Println("\nTotal Female Records:", len(Females))
	fmt.Println("Lowest Female Salary:", lowestFSalary(Females))
	maxFSalary, highestPF := MaxFSalary(Females)
	fmt.Println("Highest Female Salary:", maxFSalary)
	avgFSalary := fmt.Sprintf("%.2f", averageFSalary(Females))
	fmt.Println("Average Female Salary:", avgFSalary)
	fmt.Printf("Highest paid female:\n\tName: %v %v\n\tSSN: %v\n\tGender: %v\n\tAge: %v\n\tSalary: $%v\n",
		highestPF.fname, highestPF.lname, highestPF.ssn, highestPF.gender, highestPF.age, highestPF.salary)
	fmt.Println(strings.Repeat("-", 35))

	// Print Male Stats
	fmt.Println("Total Male Records:", len(Males))
	fmt.Println("Lowest Male Salary:", lowestMSalary(Males))
	maxMSalary, highestPM := MaxMSalary(Males)
	fmt.Println("Highest Male Salary:", maxMSalary)
	avgMSalary := fmt.Sprintf("%.2f", averageMSalary(Males))
	fmt.Println("Average Male Salary:", avgMSalary)
	fmt.Printf("Highest paid male:\n\tName: %v %v\n\tSSN: %v\n\tGender: %v\n\tAge: %v\n\tSalary: $%v\n",
		highestPM.fname, highestPM.lname, highestPM.ssn, highestPM.gender, highestPM.age, highestPM.salary)
}
