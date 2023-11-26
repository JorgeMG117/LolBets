package com.example.lolbets.model

import kotlinx.serialization.Serializable

@Serializable
data class ActiveBets (
    val bet: Bet,
    val team1: String,
    val team2: String,
    val league: String,
    val completed: Int,
)