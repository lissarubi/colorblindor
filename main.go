package main

import(
"fmt"
"strings"
"os"
"strconv"
"os/exec"
)

func execute(command string) string {
  out, err := exec.Command("bash", "-c", command).Output()

  if err != nil {}

  output := string(out[:])
  return output
}

func generateColorblindType(colorblindType string, filesSplited []string){
    simpleType := strings.Split(colorblindType, "")

    var outputFiles string
    filesSplitedLen := len(filesSplited) - 3

    for i := 0; i < filesSplitedLen; i++{

        file := filesSplited[i]
        iString := strconv.Itoa(i)

        execute("daltonize -s -t=" + simpleType[0] + " tmpColorblinder/" + file + " tmpColorblinder/" + colorblindType + "/" + iString + ".png")

        outputFiles += "tmpColorblinder/" + colorblindType + "/" + iString + ".png "
        fmt.Println(colorblindType, ":", i + 1, "° Page.")
    }

    execute("img2pdf " + outputFiles + "-o " + colorblindType + ".pdf")

}

func generateDaltonize(filesSplited []string) {
    generateColorblindType("deltan", filesSplited)
    generateColorblindType("protan", filesSplited)
    generateColorblindType("tritan", filesSplited)
}

func main() {
    if (len(os.Args) < 2){
        panic("Please, specify the PDF file.")
    }

    execute("[ -d \"tmpColorblinder\" ] && rm -rf tmpColorblinder")
    execute("mkdir tmpColorblinder tmpColorblinder/deltan tmpColorblinder/protan tmpColorblinder/tritan")
    fmt.Println("Converting PDF to images...")

    execute("convert -density 400 \"" + os.Args[1] + "\" -quality 100 \"tmpColorblinder/capture.png\"")

    fmt.Println("PDF converted to images.")

    files := execute("ls tmpColorblinder")
    filesSplited := strings.Fields(files)
    generateDaltonize(filesSplited)
    
    execute("rm -rf tmpColorblinder")
    fmt.Println("Finish! Let's check the files deltan.pdf, protan.pdf and tritan.pdf")
}
