package com.example.lolbets.data

import com.example.lolbets.R
import com.example.lolbets.model.Game
import com.example.lolbets.model.League
import com.example.lolbets.model.Team

data class BetUiState(
    val game: Game = Game(
        Team("Fnatic", "https://www.example.com/image.png"), Team("Fnatic","https://www.example.com/image.png"), League(
            R.string.league_name_lec, R.drawable.lec), "10 de junio", 100, 100),
    val teamChoice: Int = 0,
)