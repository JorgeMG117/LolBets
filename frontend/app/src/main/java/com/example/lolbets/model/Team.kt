package com.example.lolbets.model

import androidx.annotation.DrawableRes
import androidx.annotation.StringRes
import kotlinx.serialization.Serializable

@Serializable
data class TeamApi(
    val name: String,
    val code: String,
    val image: String,
)
/*data class Team(
    @StringRes val stringResourceId: Int,
    @DrawableRes val imageResourceId: Int
)*/
data class Team(
    val name: String,
    val image: String
)
