import Foundation

var result: Int = 0
let fileContents = try String(contentsOfFile: "Sources/input.txt", encoding: .utf8)
let lines = fileContents.split(separator: "\n")

for line in lines {
  let firstDigit: String? = findFirstDigit(in: String(line))
  let lastDigit: String? = findLastDigit(in: String(line))

  result += Int("\(firstDigit!)\(lastDigit!)")!
}

print(result)

func findFirstDigit(in line: String) -> String? {
  let pattern = try! Regex<Substring>("[0-9]+")
  let matches = try? pattern.firstMatch(in: line)
  if let match = matches?.first {
    return String(match)
  }
  return nil
}

func findLastDigit(in line: String) -> String? {
  let pattern = try! Regex<Substring>("[0-9]+")
  let matches = try? pattern.firstMatch(in: String(line.reversed()))
  if let match = matches?.first {
    return String(match)
  }
  return nil
}
