package com.example.lolbets.model

import androidx.annotation.DrawableRes
import androidx.annotation.StringRes

data class User(
    @StringRes val userNameResourceId: Int,
    @DrawableRes val profilePictureResourceId: Int,
    val coins: Int,

)
