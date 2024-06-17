package main

import (
        "fmt"
        "sort"
        "log"
        "flag"
        "strings"
        "maps"
        "bufio"
        "os"

        tp "textprocessing"
)

func permute(words []string) []string {
    if len(words) == 0 {
        return []string{""}
    }

    var result []string
    for i := 0; i < len(words); i++ {
        word := words[i]
        remaining := append([]string{}, words[:i]...)
        remaining = append(remaining, words[i+1:]...)
        perms := permute(remaining)
        for _, perm := range perms {
            if perm == "" {
                result = append(result, word)
            } else {
                result = append(result, word+" "+perm)
            }
        }
    }

    return result
}

func generateCombinations(input string) map[string]bool {
        words := strings.Fields(input)
  
        combinations := make(map[string]bool)

        if len(words) > 7 {
                return combinations
        }

        lines := permute(words[:len(words)-1])

        for m := 0; m < len(lines);m++ {
            lineWords := strings.Fields(lines[m])
            for i := 0; i < len(lineWords); i++ {
                for j := i + 1; j < len(lineWords); j++ {
                        joined := lineWords[:i]
                        sort.Strings(joined)
                        secJoined := lineWords[i:j+1]
                        sort.Strings(secJoined)
                        ccmb := strings.Fields(strings.TrimPrefix(strings.Join(joined, "~") + " ", " ") + strings.TrimPrefix(strings.Join(secJoined, "~")+ " ", " ") + strings.TrimPrefix(strings.Join(lineWords[j+1:len(lineWords)]," ") + " ", " "))
                        sort.Strings(ccmb)
                        combinations[strings.Join(ccmb, " ") + " " + words[len(words)-1]] = true
                }
             }
        }

        return combinations
}

func main() {
        NormalizerFilename := flag.String(
                "normalizer dictionary", "words-exactmatch.latest_22.04.24.csv.gz", "normalizer file")
        normalizedDictionary := make(map[string]bool)
        flag.Parse()


        // Write output file
        outfile, err := os.Create(fmt.Sprintf("mqData%s.csv", "6m"))
        if err != nil {
                log.Fatal(err)
        }
        defer outfile.Close()

        writer := bufio.NewWriter(outfile)
        defer writer.Flush()
        tpr, err := tp.New(
                tp.LoadNormalizerFromFilename(*NormalizerFilename),
        )

        qf, err := os.Open("6m_random_q_to_nm.csv")
        if err != nil {
                panic(err)
        }
        defer qf.Close()

        scanner := bufio.NewScanner(qf)

        i := 0
        for scanner.Scan() {
                norm := strings.ReplaceAll(scanner.Text(), ",", " ")
                words := strings.Fields(norm)

                lastWord := words[len(words)-1]
                previous := tpr.Normalization(strings.Join(words[:len(words)-1], " "))
                normalizedDictionary[strings.ToLower(previous + " " + lastWord)] = true
                i++
        }

        output := make(map[string]bool)
        i = 0
        for dict := range normalizedDictionary {
                if i % 10000 == 0{
                        fmt.Println(i, "/", len(normalizedDictionary))
                }
                i++
                output[dict] = true
                maps.Copy(output, generateCombinations(dict))
        }

        for out := range output {
                writer.WriteString(out + "\n")
                writer.Flush()
        }

        fmt.Println(len(output))
}
