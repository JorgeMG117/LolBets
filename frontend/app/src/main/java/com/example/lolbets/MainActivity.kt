package com.example.lolbets

import android.R
import android.content.ContentValues.TAG
import android.content.IntentSender.SendIntentException
import android.os.Bundle
import android.util.Log
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.compose.ui.Modifier
//import com.example.lolbets.network.KtorRealtimeMessagingClient
import com.example.lolbets.ui.theme.LolBetsTheme
import com.google.android.gms.auth.api.identity.GetSignInIntentRequest
import com.google.android.gms.auth.api.identity.Identity
import com.google.android.gms.auth.api.signin.GoogleSignIn
import com.google.android.gms.auth.api.signin.GoogleSignInClient
import com.google.android.gms.auth.api.signin.GoogleSignInOptions



class MainActivity : ComponentActivity() {

    private fun getGoogleLoginAuth(): GoogleSignInClient {
        val gso = GoogleSignInOptions.Builder(GoogleSignInOptions.DEFAULT_SIGN_IN)
            .requestEmail()
            /*.requestIdToken("")
            .requestId()
            .requestProfile()*/
            .build()
        return GoogleSignIn.getClient(this, gso)
    }

    private val REQUEST_CODE_GOOGLE_SIGN_IN = 1 /* unique request id */

    private fun signIn() {
        val request = GetSignInIntentRequest.builder()
            .setServerClientId("")
            .build()
        Identity.getSignInClient(this)
            .getSignInIntent(request)
            .addOnSuccessListener { result ->
                try {
                    startIntentSenderForResult(
                        result.getIntentSender(),
                        REQUEST_CODE_GOOGLE_SIGN_IN,  /* fillInIntent= */
                        null,  /* flagsMask= */
                        0,  /* flagsValue= */
                        0,  /* extraFlags= */
                        0,  /* options= */
                        null
                    )
                } catch (e: SendIntentException) {
                    Log.e(TAG, "Google Sign-in failed")
                }
            }
            .addOnFailureListener { e -> Log.e(TAG, "Google Sign-in failed", e) }
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            LolBetsTheme {
                // A surface container using the 'background' color from the theme
                Surface(
                    modifier = Modifier.fillMaxSize(),
                    color = MaterialTheme.colorScheme.background
                ) {
                    //val glogauth = getGoogleLoginAuth()
                    val gso = GoogleSignInOptions.Builder(GoogleSignInOptions.DEFAULT_SIGN_IN)
                        .requestEmail()
                        .build()
                    val mGoogleSignInClient = GoogleSignIn.getClient(this, gso);
                    //signIn()
                    LolBetsApp(mGoogleSignInClient)
                }
            }
        }
    }
}