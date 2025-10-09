package main

import (
	"fmt"
)

func main() {
	PreClassExercise()

	HemoglobinUPGMA()

	SARS2UPGMA()
}

func PreClassExercise() {
	fmt.Println("Read in Pre-class Exercise matrix.")

	speciesNames, mtx := ReadMatrixFromFile("Data/PreClassExercise/distance.mtx")

	fmt.Println("Starting UPGMA.")

	//generate UPGMA tree
	preclassExerciseTree := UPGMA(mtx, speciesNames)

	fmt.Println("UPGMA tree built. Writing to file.")

	WriteNewickToFile(preclassExerciseTree, "Output/PreClassExercise", "exercise.tre")

	fmt.Println("Tree written to file.")
}

func HemoglobinUPGMA() {
	fmt.Println("Read in Hemoglobin alpha subunit 1 matrix.")

	speciesNames, mtx := ReadMatrixFromFile("Data/HBA1/hemoglobin.mtx")

	fmt.Println("Starting UPGMA.")

	//generate UPGMA tree
	hemoglobinTree := UPGMA(mtx, speciesNames)

	fmt.Println("UPGMA tree built. Writing to file.")

	WriteNewickToFile(hemoglobinTree, "Output/HBA1", "hba1.tre")

	fmt.Println("Tree written to file.")
}

func SARS2UPGMA() {
	fmt.Println("Read in SARS-CoV-2 matrix.")

	genomeLabels, mtx := ReadMatrixFromFile("Data/UK-SARS-CoV-2/SARS_spike.mtx")

	fmt.Println("Matrix read!")

	fmt.Println("Generating UPGMA tree.")

	//generate UPGMA tree
	sarsTree := UPGMA(mtx, genomeLabels)

	fmt.Println("UPGMA tree built! Writing to file.")

	WriteNewickToFile(sarsTree, "Output/UK-SARS-CoV-2", "sars-cov-2.tre")

	fmt.Println("Tree written to file.")
}
