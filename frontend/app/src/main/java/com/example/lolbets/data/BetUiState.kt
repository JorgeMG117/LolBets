package com.example.lolbets.data

import com.example.lolbets.R
import com.example.lolbets.model.Game
import com.example.lolbets.model.League
import com.example.lolbets.model.Team

data class BetUiState(
    val game: Game = Game(
        Team(R.string.team_name_astralis, R.drawable.astralis), Team(R.string.team_name_fnatic, R.drawable.fnatic), League(
            R.string.league_name_lec, R.drawable.lec), "10 de junio", 100, 100),
    val teamChoice: Int = 0,
)