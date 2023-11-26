package com.example.lolbets.model

import kotlinx.serialization.Serializable

@Serializable
data class Bet(
    val value: Int,
    val team: Boolean,
    val userId: Int,
    val gameId: Int,
    val odds: Double,
)