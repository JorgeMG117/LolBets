package com.example.lolbets.model

import androidx.compose.ui.graphics.vector.ImageVector

data class BottomNavItem(
    val name: String,
    val onButtonClicked: () -> Unit,
    val icon: ImageVector,
)

