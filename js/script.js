// Copyright (c) 2023 Charlotte Jhu All rights reserved
//
// Created by: Charlotte Jhu
// Created on: May 2023
// This file contains the JS functions for index.html

'use strict'

function myButtonClicked() {
// This function does calculates the hailstone sequence
// input
  let positiveInteger = parseInt(document.getElementById("positive-integer").value)

// process
  while (positiveInteger > 1) {
    if (positiveInteger%2 == 0) {
      positiveInteger = positiveInteger / 2
      document.getElementById("answer").innerHTML = positiveInteger
    } else {
      positiveInteger = positiveInteger * 3 + 1
      document.getElementById("answer").innerHTML = positiveInteger
    }
  }

// output
  document.getElementById("answer").innerHTML = positiveInteger
}
