package com.example.lolbets.ui

import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.lolbets.model.GameApi
import com.example.lolbets.network.LolApi
import kotlinx.coroutines.launch
import retrofit2.HttpException
import java.io.IOException


sealed interface GameUiState {
    data class Success(val games: String) : GameUiState
    object Error : GameUiState
    object Loading : GameUiState
}

class GamesViewModel : ViewModel() {
    var gameUiState: GameUiState by mutableStateOf(GameUiState.Loading)
        private set


    init {
        getMarsPhotos()
    }

    fun getMarsPhotos() {
        viewModelScope.launch {
            gameUiState = GameUiState.Loading
            gameUiState = try {
                val listResult = LolApi.retrofitService.getGames()
                GameUiState.Success(
                    "Success: ${listResult.size} Mars photos retrieved"
                )
            } catch (e: IOException) {
                GameUiState.Error
            } catch (e: HttpException) {
                GameUiState.Error
            }
        }
    }
}
