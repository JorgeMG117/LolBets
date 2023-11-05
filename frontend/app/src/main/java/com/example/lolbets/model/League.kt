package com.example.lolbets.model

import androidx.annotation.DrawableRes
import androidx.annotation.StringRes
import kotlinx.serialization.Serializable

@Serializable
data class LeagueApi(
    val id: String,
    val name: String,
    val region: String,
    val image: String,
)

data class League(
    @StringRes val stringResourceId: Int,
    @DrawableRes val imageResourceId: Int
)
