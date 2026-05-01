package model

type State struct {
  X, Y int
  NextNumber int // angka berikutnya yang harus dilewati
  Cost int       // g(n)
}