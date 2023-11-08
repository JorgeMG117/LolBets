package com.example.lolbets.model

import androidx.annotation.DrawableRes
import androidx.annotation.StringRes
import kotlinx.serialization.Serializable

@Serializable
data class Team(
    val name: String,
    val code: String,
    val image: String
)
