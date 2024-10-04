package models

import (
  "gorm.io/gorm"
)

type Ideas struct {
  gorm.Model

  id string 'json:"id" gorm:"primaryKey"'
  Idea string 'json:"idea" gorm:"not null"'
  Judge1Score uint 'json:"jg1score"'
  Judge2Score uint 'json:"jg2score"'
  Judge3Score uint 'json:"jg3score"'
  AvgScore float32 'json:"avgscore"'
}
