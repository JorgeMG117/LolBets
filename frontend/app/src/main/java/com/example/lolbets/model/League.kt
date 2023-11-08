package com.example.lolbets.model

import androidx.annotation.DrawableRes
import androidx.annotation.StringRes
import kotlinx.serialization.Serializable

@Serializable
data class League(
    val id: String,
    val name: String,
    val region: String,
    val image: String,
)

