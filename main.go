/****************************************************************
Joust
Author: Bryan Saucedo-Mondragon
Date Completed: 10/27/24
Description:create a program, that is meant to joust two knight, use flag commands. A whole bunch of functions, and member functions.
****************************************************************/
package main

import (
  "flag"
  "fmt"
  "math/rand"
)

var (
  Knight1Name        string
  Knight1Stamina     int
  Knight1WeaponType  string
  Knight1HitChance   int
  Knight1StaminaCost int
  Knight2Name        string
  Knight2Stamina     int
  Knight2WeaponType  string
  Knight2HitChance   int
  Knight2StaminaCost int
)


type Knight struct {
  Name    string
  Stamina int
  Mounted bool
  Weapon  Weapon
}

func NewKnight(name string, stamina int, weapon Weapon) Knight {
  return Knight{
    Name:    name,
    Stamina: stamina,
    Mounted: true,
    Weapon:  weapon,
  }
}

func (k Knight) GetName() string {
  return k.Name
}
func (k Knight) GetWeapon() Weapon {
  return k.Weapon
}

func (k Knight) GetStamina() int {
  return k.Stamina
}

func (k Knight) GetMounted() bool {
  return k.Mounted
}

func (k *Knight) SetMounted(mounted bool) {
  k.Mounted = mounted
}

func (k *Knight) Joust() bool {
  hit := rand.Intn(100) < k.Weapon.HitChance
  k.Stamina -= k.Weapon.StaminaCost
  return hit
}

func (k Knight) DisplayStats() {
  var mountedStatus string
  if k.Mounted {
    mountedStatus = "is still on horse."
  } else {
    mountedStatus = "has been knocked off of the horse."
  }
  var staminaStatus string
  if k.Stamina > 0 {
    staminaStatus = "is not exhausted"
  } else {
    staminaStatus = "is exhausted"
  }
  if k.Stamina > 0 {
    fmt.Print(k.Name, " ", staminaStatus, " (stamina=", k.Stamina, ") and ", mountedStatus, "\n")
    fmt.Print(k.Name, " is using: ", k.Weapon.Type, " that requires ", k.Weapon.StaminaCost, " stamina and has a ", k.Weapon.HitChance, "% chance to hit.\n")
  } else {
    fmt.Print(k.Name, " ", staminaStatus, " and ", mountedStatus, "\n")
    fmt.Print(k.Name, " is using: ", k.Weapon.Type, " that requires ", k.Weapon.StaminaCost, " stamina and has a ", k.Weapon.HitChance, "% chance to hit.\n")
  }
}

type Weapon struct {
  StaminaCost int
  HitChance   int
  Type        string
}

func NewWeapon(staminaCost int, hitChance int, weaponType string) Weapon {
  return Weapon{
    StaminaCost: staminaCost,
    HitChance:   hitChance,
    Type:        weaponType,
  }
}

func (w Weapon) GetStaminaCost() int {
  return w.StaminaCost
}
func (w Weapon) GetHitChance() int {
  return w.HitChance
}
func (w Weapon) GetType() string {
  return w.Type
}
func (w Weapon) Swing() bool {
  if rand.Intn(100) <= w.HitChance {
    return true
  } else {
    return false
  }
}


func init() {
  fmt.Println("Jousting game is initializing...")
}


func main() {
  flag.StringVar(&Knight1Name, "Knight1Name", "King Arthur", "Name of Knight 1")
  flag.IntVar(&Knight1Stamina, "Knight1Stamina", 50, "Stamina of Knight 1")
  flag.StringVar(&Knight1WeaponType, "Knight1WeaponType", "Excalibur", "Weapon type of Knight 1")
  flag.IntVar(&Knight1HitChance, "Knight1HitChance", 15, "Hit chance of Knight 1")
  flag.IntVar(&Knight1StaminaCost, "Knight1StaminaCost", 10, "Stamina cost of Knight 1")

  flag.StringVar(&Knight2Name, "Knight2Name", "Black Knight", "Name of Knight 2")
  flag.IntVar(&Knight2Stamina, "Knight2Stamina", 40, "Stamina of Knight 2")
  flag.StringVar(&Knight2WeaponType, "Knight2WeaponType", "Longsword", "Weapon type of Knight 2")
  flag.IntVar(&Knight2HitChance, "Knight2HitChance", 10, "Hit chance of Knight 2")
  flag.IntVar(&Knight2StaminaCost, "Knight2StaminaCost", 5, "Stamina cost of Knight 2")

  
  flag.Parse()


  Knight1Weapon := NewWeapon(Knight1StaminaCost, Knight1HitChance, Knight1WeaponType)
  Knight2Weapon := NewWeapon(Knight2StaminaCost, Knight2HitChance, Knight2WeaponType)

  knight1 := NewKnight(Knight1Name, Knight1Stamina, Knight1Weapon)
  knight2 := NewKnight(Knight2Name, Knight2Stamina, Knight2Weapon)

  
  for knight1.GetMounted() && knight2.GetMounted() && knight1.GetStamina() > 0 && knight2.GetStamina() > 0 {
    hitKnight2 := knight1.Joust()
    if hitKnight2 {
      knight2.SetMounted(false)
    }

    hitKnight1 := knight2.Joust()
    if hitKnight1 {
      knight1.SetMounted(false)
    }

    knight1.DisplayStats()
    knight2.DisplayStats()
    fmt.Print("\n")
  }

  // Determine the winner
  if (!knight1.GetMounted() || knight1.GetStamina() <= 0) && (!knight2.GetMounted() || knight2.GetStamina() <= 0) {
    fmt.Println("It's a draw!")
  } else if !knight1.GetMounted() || knight1.GetStamina() <= 0 {
    fmt.Print(knight2.Name, " wins!\n")
  } else {
    fmt.Print(knight1.Name, " wins!\n")
  }
}
