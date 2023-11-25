package com.example.lolbets.data

import com.example.lolbets.R
import com.example.lolbets.model.Game
import com.example.lolbets.model.League
import com.example.lolbets.model.Team

data class BetUiState(
    val game: Game = Game(
        5,
        Team("Fnatic", "Fnatic", "https://www.example.com/image.png"),
        Team("Fnatic", "Fnatic","https://www.example.com/image.png"),
        League("0", "LEC", "EMEA", "https://www.example.com/image.png"),
        "10 de junio", 100, 100, 0, "",""),
    val teamChoice: Int = 0,
    //val isConnected: Boolean = false
)