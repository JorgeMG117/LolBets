package com.example.lolbets.model

import androidx.annotation.DrawableRes
import androidx.annotation.StringRes

data class Team(
    @StringRes val stringResourceId: Int,
    @DrawableRes val imageResourceId: Int
)
