///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// LICENSE: GNU General Public License, version 2 (GPLv2)
// Copyright 2016, Charlie J. Smotherman
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License v2
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	ampgolib "github.com/cjsmocjsmo/ampgolibmod"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func initialArtistInfo(w http.ResponseWriter, r *http.Request) {
	AV := ampgolib.GetInitArtistInfo()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AV)
	log.Println("Initial Artist Info Complete")
	return
}

func initialalbumInfoHandler(w http.ResponseWriter, r *http.Request) {
	GIAI := ampgolib.GetInitAlbumInfo()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GIAI)
	return
}

func initialsongInfoHandler(w http.ResponseWriter, r *http.Request) {
	ISI := ampgolib.GetInitSongInfo()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ISI)
	return
}

func artistPageHandler(w http.ResponseWriter, r *http.Request) {
	ArtA := ampgolib.ArtistAlpha()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ArtA)
}

func albumPageHandler(w http.ResponseWriter, r *http.Request) {
	AlbA := ampgolib.AlbumAlpha()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AlbA)
}

func titlePageHandler(w http.ResponseWriter, r *http.Request) {
	TA := ampgolib.TitleAlpha()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TA)
}

func imageSongsForAlbumHandler(w http.ResponseWriter, r *http.Request) {
	qu := r.URL.Query().Get("selected")
	foo := ampgolib.GImageSongForAlbum(qu)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foo)
	return
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	ST := ampgolib.GStats()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ST)
}

// func randomPicsHandler(w http.ResponseWriter, r *http.Request) {
// 	RandomPics := ampgolib.RPics()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(RandomPics)
// }

// func ramdomAlbumPicPlaySongHandler(w http.ResponseWriter, r *http.Request) {
// 	qu := r.URL.Query().Get("sid")
// 	rapp := ampgolib.RamdomAlbPicPlay(qu)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(rapp)
// }

// func pathArtHandler(w http.ResponseWriter, r *http.Request) {
// 	qu := r.URL.Query().Get("selected")
// 	pa := ampgolib.PathArt(qu)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(pa)
// }

func songInfoHandler(w http.ResponseWriter, r *http.Request) {
	qu := r.URL.Query().Get("selected")
	si := ampgolib.SongInfo(qu)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(si)
}

func albumInfoHandler(w http.ResponseWriter, r *http.Request) {
	qu := r.URL.Query().Get("selected")
	ai := ampgolib.AlbumInfo(qu)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ai)
}

func artistInfoHandler(w http.ResponseWriter, r *http.Request) {
	qu := r.URL.Query().Get("selected")
	arti := ampgolib.ArtistInfo(qu)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(arti)
}

// func songSearchHandler(w http.ResponseWriter, r *http.Request) {
// 	qu := r.URL.Query().Get("searchval")
// 	artS := ampgolib.SongSearch(qu)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(artS)
// }

// func albumSearchHandler(w http.ResponseWriter, r *http.Request) {
// 	qu := r.URL.Query().Get("albsearchval")
// 	albS := ampgolib.AlbumSearch(qu)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(albS)
// }

// func artistSearchHandler(w http.ResponseWriter, r *http.Request) {
// 	qu := r.URL.Query().Get("artsearchval")
// 	artS := ampgolib.ArtistSearch(qu)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(artS)
// }

// func allPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	plc := ampgolib.PlaylistCheck()
// 	if plc != 1 {
// 		json.NewEncoder(w).Encode("Please create a playlist")
// 	} else {
// 		allpls := ampgolib.AllPlayLists()
// 		json.NewEncoder(w).Encode(allpls)
// 	}
// }

// func addPlayListNameToDBHandler(w http.ResponseWriter, r *http.Request) {
// 	qu := r.URL.Query().Get("playlistname")
// 	aplntdb := ampgolib.AddPlayListNameToDB(qu)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(aplntdb)
// }

// func addSongsToPlistDBHandler(w http.ResponseWriter, r *http.Request) {
// 	sn := r.URL.Query().Get("songname")
// 	sid := r.URL.Query().Get("songid")
// 	plid := r.URL.Query().Get("playlistid")
// 	ampgolib.AddSongsToPlistDB(sn, sid, plid)
// }

// func allPlaylistSongsFromDBHandler(w http.ResponseWriter, r *http.Request) {
// 	plid := r.URL.Query().Get("playlistid")
// 	soho := ampgolib.AllPlaylistSongsFromDB(plid)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(soho)
// }

// func createPlayerPlaylistHandler(w http.ResponseWriter, r *http.Request) {
// 	plid := r.URL.Query().Get("playlistid")
// 	apsf := ampgolib.CreatePlayerPlaylist(plid)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(apsf)
// }

// func addRandomPlaylistHandler(w http.ResponseWriter, r *http.Request) {
// 	plname := r.URL.Query().Get("playlistname")
// 	plcount := r.URL.Query().Get("playlistcount")
// 	rpl := ampgolib.AddRandomPlaylist(plname, plcount)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(rpl)
// }

// func deletePlaylistFromDBHandler(w http.ResponseWriter, r *http.Request) {
// 	plid := r.URL.Query().Get("playlistid")
// 	dpl := ampgolib.DeletePlaylistFromDB(plid)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(dpl)
// }

// func deleteSongFromPlaylistHandler(w http.ResponseWriter, r *http.Request) {
// 	plname := r.URL.Query().Get("playlistname")
// 	songid := r.URL.Query().Get("delsongid")
// 	dsfp := ampgolib.DeleteSongFromPlaylist(plname, songid)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(dsfp)
// }

// func setUpHandler(w http.ResponseWriter, r *http.Request) {
// 	ampgolib.SetUp()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode("Setup Complete")
// }

// func setUpHandler(w http.ResponseWriter, r *http.Request) {
// 	ampgolib.SetUp()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode("Setup Complete")
// }

func init() {
	ampgolib.SetUpCheck()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/InitialArtistInfo", initialArtistInfo)
	r.HandleFunc("/InitialAlbumInfo", initialalbumInfoHandler)
	r.HandleFunc("/InitialSongInfo", initialsongInfoHandler)

	r.HandleFunc("/ArtistAlpha", artistPageHandler)
	r.HandleFunc("/AlbumAlpha", albumPageHandler)
	r.HandleFunc("/TitleAlpha", titlePageHandler)

	r.HandleFunc("/ArtistInfo", artistInfoHandler)
	r.HandleFunc("/AlbumInfo", albumInfoHandler)
	r.HandleFunc("/SongInfo", songInfoHandler)

	r.HandleFunc("/ImageSongsForAlbum", imageSongsForAlbumHandler)

	// r.HandleFunc("/RandomPics", randomPicsHandler)
	// r.HandleFunc("/RamdomAlbumPicPlaySong", ramdomAlbumPicPlaySongHandler)
	r.HandleFunc("/Stats", statsHandler)
	// r.HandleFunc("/PathArt", pathArtHandler)
	// r.HandleFunc("/SongSearch", songSearchHandler)
	// r.HandleFunc("/AlbumSearch", albumSearchHandler)
	// r.HandleFunc("/ArtistSearch", artistSearchHandler)
	// r.HandleFunc("/AllPlaylists", allPlaylistsHandler)
	// r.HandleFunc("/AddPlayListNameToDB", addPlayListNameToDBHandler)
	// r.HandleFunc("/AddSongsToPlistDB", addSongsToPlistDBHandler)
	// r.HandleFunc("/AllPlaylistSongsFromDB", allPlaylistSongsFromDBHandler)
	// r.HandleFunc("/CreatePlayerPlaylist", createPlayerPlaylistHandler)
	// r.HandleFunc("/AddRandomPlaylist", addRandomPlaylistHandler)
	// r.HandleFunc("/DeletePlaylistFromDB", deletePlaylistFromDBHandler)
	// r.HandleFunc("/DeleteSongFromPlaylist", deleteSongFromPlaylistHandler)
	// r.HandleFunc("/SetUp", setUpHandler)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))
	http.ListenAndServe(":8080", (r))
	// 		csrf.Protect([]byte(key), csrf.Secure(false))(r))
}
