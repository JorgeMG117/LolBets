package com.example.lolbets.viewmodel

import androidx.lifecycle.ViewModel
import com.example.lolbets.R
import com.example.lolbets.data.BetUiState
import com.example.lolbets.model.Game
import com.example.lolbets.model.User
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update

class UserViewModel : ViewModel() {

    private val _userState = MutableStateFlow(User(R.string.team_name_koi, R.drawable.koi, 1000))
    val userState: StateFlow<User> = _userState.asStateFlow()

    fun placeBet(value: Int) {
        // TODO Check if value is negative

        _userState.update { currentState ->
            currentState.copy(
                coins = currentState.coins - value
            )
        }
    }
}