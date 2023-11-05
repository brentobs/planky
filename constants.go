package main

var languages = []string{"Python", "NodeJS", "Java", "Go"}
var pythonTypes = []string{"FastAPI", "Django", "CLI"}
var nodeTypes = []string{"Express", "NestJS", "CLI"}
var javaTypes = []string{"Spring", "Bla", "CLI"}
var goTypes = []string{"a", "b", "CLI"}
var languageTypes = map[string][]string{
    "Python": {"FastAPI", "Django", "CLI"},
    "NodeJS": {"Express", "NestJS", "CLI"},
    "Java":   {"Spring", "Bla", "CLI"},
    "Go":     {"a", "b", "CLI"},
}
