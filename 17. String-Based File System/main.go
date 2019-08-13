package main

import (
	"fmt"
	"strings"
)

// system by a string in the following manner:
// The string "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext" represents:
// dir
//     subdir1
//     subdir2
//         file.ext
// The directory dir contains an empty sub-directory subdir1 and a sub-directory subdir2 containing a file file.ext.
// The string "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext" represents:
// dir
//     subdir1
//         file1.ext
//         subsubdir1
//     subdir2
//         subsubdir2
//             file2.ext
// The directory dir contains two sub-directories subdir1 and subdir2. subdir1 contains a file file1.ext and an empty second-level sub-directory subsubdir1. subdir2 contains a second-level sub-directory subsubdir2 containing a file file2.ext.
// We are interested in finding the longest (number of characters) absolute path to a file within our file system. For example, in the second example above, the longest absolute path is "dir/subdir2/subsubdir2/file2.ext", and its length is 32 (not including the double quotes).

func main() {

	fileSystem := "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"

	fmt.Println(GetLongestPathSize(fileSystem))
}

func GetLongestPathSize(filePath string) int {

	splitPaths := strings.Split(filePath, "\n\t")

	basePath := splitPaths[0]
	subPaths := splitPaths[1:]

	longestSubPath := Helper(subPaths)

	longestPath := basePath + "/" + longestSubPath

	fmt.Println(longestPath)

	return len(longestPath)
}

func Helper(filePaths []string) string {

	longestPath := ""

	for idx, filePath := range filePaths {

		if strings.HasPrefix(filePath, "\t") {
			longestSubPath := Helper(filePath[2:])			
		}
		subPaths := strings.Split(filePath, "\t")

		for
	}

	if len(subPaths) == 1 && !strings.Contains(subPaths[0], "\n\t") {
		return subPaths[0]
	}

	basePath := subPaths[0]
	largestPath := ""

	for i := 1; i < len(subPaths); i++ {

		largestSubPath := Helper(subPaths[i], depth+1)

		if len(largestSubPath) > len(largestPath) {
			largestPath = largestSubPath
		}
	}

	return basePath + largestPath
}

// func Helper(filePath string, depth int) string {

// 	filePathSeperator := GetFileSeparator(depth)

// 	firstSeperator := strings.Index(filePath, filePathSeperator)
// 	if firstSeperator == -1 {
// 		return filePath
// 	}

// 	basePath := filePath[:firstSeperator]
// 	largestPath := ""

// 	for {
// 		if firstSeperator == -1 {
// 			break
// 		}

// 		filePath = filePath[firstSeperator+len(filePathSeperator):]
// 		// fmt.Print(filePath)

// 		pathSize := Helper(filePath, depth+1)

// 		if pathSize > largestPath {
// 			largestPath = pathSize
// 		}

// 		firstSeperator = strings.Index(filePath, filePathSeperator)
// 	}

// 	return basePath + "/" + largestPath
// }

func GetFileSeparator(depth int) string {
	seperator := "\n"

	for i := 0; i <= depth; i++ {
		seperator += "\t"
	}

	return seperator
}
