package com.example.lolbets.viewmodel

import androidx.lifecycle.ViewModel
import com.example.lolbets.data.BetUiState
import com.example.lolbets.model.Game
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update

class FocusedGameViewModel : ViewModel() {

    private val _uiState = MutableStateFlow(BetUiState())
    val uiState: StateFlow<BetUiState> = _uiState.asStateFlow()

    fun setGame(game: Game) {
        _uiState.update { currentState ->
            currentState.copy(
                game = game
            )
        }
    }
}