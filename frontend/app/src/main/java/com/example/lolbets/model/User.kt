package com.example.lolbets.model

import androidx.annotation.DrawableRes
import androidx.annotation.StringRes
import kotlinx.serialization.Serializable

@Serializable
data class UserApi(
    val id: Int,
    val name: String,
    val coins: Int,
)

data class User(
    val id: Int,
    val googleId: String,
    val coins: Int,
    val username: String?,
    val profilePictureUrl: String?,
)
