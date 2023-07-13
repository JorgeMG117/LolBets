package com.example.lolbets.model

import androidx.annotation.DrawableRes
import androidx.annotation.StringRes

data class Game(
    val team1: Team,
    val team2: Team,
    val league: League,
    val date: String,
    val betsTeam1: Int,
    val betsTeam2: Int,
)
