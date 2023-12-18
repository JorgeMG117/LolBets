package com.example.lolbets.viewmodel

import android.util.Log
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.lolbets.data.BetUiState
import com.example.lolbets.model.Bet
import com.example.lolbets.model.Game
//import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.SharingStarted
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.catch
import kotlinx.coroutines.flow.onEach
import kotlinx.coroutines.flow.onStart
import kotlinx.coroutines.flow.stateIn
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch
import java.net.ConnectException
import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import com.example.lolbets.model.ActiveBets
import com.example.lolbets.model.League
import com.example.lolbets.model.Team
import com.example.lolbets.network.BetResponse
import com.example.lolbets.network.MyWebSocketListener

import kotlinx.coroutines.Dispatchers
import kotlinx.serialization.json.Json
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.Response
import okhttp3.WebSocket
import okhttp3.WebSocketListener
import okio.ByteString
import java.util.concurrent.TimeUnit

class FocusedGameViewModel(private val onFocusedOut: (Game) -> Unit) : ViewModel() {

    private val _uiState = MutableStateFlow(BetUiState())
    val uiState: StateFlow<BetUiState> = _uiState.asStateFlow()

    fun setGame(game: Game) {
        _uiState.update { currentState ->
            currentState.copy(
                game = game
            )
        }
    }

    fun setUserId(userId: Int) {
        _uiState.update { currentState ->
            currentState.copy(
                userId = userId
            )
        }
    }

    fun updateGameBets(message: BetResponse) {
        if(uiState.value.game.betsTeam1 != message.bets1 || uiState.value.game.betsTeam2 != message.bets2) {
            //uiState.value.game.betsTeam1 = message.bets1
            //game = Game(uiState.value.game.betsTeam1)
            val updatedGame = Game(
                uiState.value.game.id,
                uiState.value.game.team1,
                uiState.value.game.team2,
                uiState.value.game.league,
                uiState.value.game.date,
                message.bets1,
                message.bets2,
                uiState.value.game.completed,
                uiState.value.game.blockName,
                uiState.value.game.strategy,
            )
            _uiState.update { currentState ->
                currentState.copy(
                    game = updatedGame
                )
            }
        }
    }

    fun betSucceed() {
        //Cambiar val a true
        _uiState.update { currentState ->
            currentState.copy(
                betSucceed = true
            )
        }
    }

    fun restartBet() {
        _uiState.update { currentState ->
            currentState.copy(
                betSucceed = false
            )
        }
    }

    //Websockets
    private val _socketStatus = MutableLiveData(false)
    val socketStatus: LiveData<Boolean> = _socketStatus

    //private val _messages = MutableLiveData<Pair<Boolean, String>>()
    //val messages: LiveData<Pair<Boolean, String>> = _messages

    /*fun addMessage(message: Pair<Boolean, String>) = viewModelScope.launch(Dispatchers.Main) {
        if (_socketStatus.value == true) {
            _messages.value = message
        }
    }*/

    fun setStatus(status: Boolean) = viewModelScope.launch(Dispatchers.Main) {
        _socketStatus.value = status
    }

    private var webSocket: WebSocket? = null

    fun connectWebSocket(gameId: Int, onBetSuccess: (ActiveBets) -> Unit) {
        val client = OkHttpClient()

        val request = Request.Builder()
            .url("wss://lolbets-image-y5oknlrakq-no.a.run.app:443/bets?game=$gameId")
            .build()

        val listener = MyWebSocketListener(this, onBetSuccess)
        val conn = client.newWebSocket(request, listener)
        webSocket = conn
    }
    /*fun connectWebSocket() {
        val client = OkHttpClient.Builder()
            .readTimeout(20, TimeUnit.SECONDS) // Adjust the timeout as needed
            .build()


    }*/

    fun sendMessage(message: Bet) {
        _uiState.update { currentState ->
            currentState.copy(
                lastBet = message
            )
        }
        println("Sending bet")
        println(message)
        val serializer = Bet.serializer()
        webSocket?.send(Json.encodeToString(serializer, message))
    }

    // Function to disconnect the WebSocket
    fun disconnectWebSocket() {
        //onDisconnectBet: (ActiveBets) -> Unit
        // Update game odds from Games list
        onFocusedOut(uiState.value.game)

        if(uiState.value.betSucceed) {
            restartBet()
        }
        webSocket?.close(1000, "User disconnected")
    }

}