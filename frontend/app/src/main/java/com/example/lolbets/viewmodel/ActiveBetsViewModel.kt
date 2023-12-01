package com.example.lolbets.viewmodel

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.lolbets.data.BetUiState
import com.example.lolbets.model.ActiveBets
import com.example.lolbets.model.Game
import com.example.lolbets.model.League
import com.example.lolbets.model.Team
import com.example.lolbets.network.LolApi
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import retrofit2.HttpException
import retrofit2.Response
import java.io.IOException
import java.text.SimpleDateFormat
import java.time.ZonedDateTime
import java.time.format.DateTimeFormatter
import java.util.Locale


sealed interface ActiveBetsUiState {
    data class Success(val activeBets: List<ActiveBets>) : ActiveBetsUiState
    object Error : ActiveBetsUiState
    object Loading : ActiveBetsUiState
}

class ActiveBetsViewModel(private var userId: Int) : ViewModel() {
    var activeBetsUiState: ActiveBetsUiState by mutableStateOf(ActiveBetsUiState.Loading)
        private set

    init {
        getActiveBets()
    }

    private fun getActiveBets() {
        viewModelScope.launch {
            activeBetsUiState = ActiveBetsUiState.Loading
            activeBetsUiState = try {
                val activeBetsList = LolApi.retrofitService.getActiveBets(userId)

                ActiveBetsUiState.Success(
                    activeBetsList
                )
            } catch (e: IOException) {
                ActiveBetsUiState.Error
            } catch (e: HttpException) {
                ActiveBetsUiState.Error
            }
        }
    }

    fun updateActiveBets(newUser: Int) {
        println(newUser)
        userId = newUser
        getActiveBets()
    }

    fun addActiveBet(bet: ActiveBets) {
        /*when (val currentState = activeBetsUiState) {
            is ActiveBetsUiState.Success -> {
                // Get the current list of active bets and add the new bet
                val updatedList = currentState.activeBets.toMutableList().apply {
                    add(bet)
                }

                // Update the UI state with the new list
                activeBetsUiState = ActiveBetsUiState.Success(updatedList)
            }
            // Ignore adding the bet if the UI state is not Success
            else -> {
                // You can choose to handle this case differently if needed
            }
        }*/
        /*
        _uiState.update { currentState ->
            currentState.copy(
                game = game
            )
        }
         */
        //println(activeBetsUiState)
        val updatedList = (activeBetsUiState as ActiveBetsUiState.Success).activeBets.toMutableList()
        updatedList.add(0, bet)
        //println(updatedList.size)
        activeBetsUiState = ActiveBetsUiState.Success(updatedList)


    }
}
