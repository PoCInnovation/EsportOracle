"use client";
import React, { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { NavBar } from "@/components/navbar";
import { useUser } from "@/app/account/useUser";

// Single definition of AutoApplyAllUsers
type AutoApplyAllUsersProps = { eventId: string; onApplied: () => void };
const AutoApplyAllUsers: React.FC<AutoApplyAllUsersProps> = ({ eventId, onApplied }) => {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);
    const [success, setSuccess] = useState(false);

    const handleAutoApply = async () => {
        setLoading(true);
        setError(null);
        setSuccess(false);
        try {
            const res = await fetch(`/api/events/${eventId}/auto-apply-all`, { method: "POST" });
            if (!res.ok) throw new Error("Erreur lors de l'auto-apply");
            setSuccess(true);
            onApplied();
        } catch (err: any) {
            setError(err instanceof Error ? err.message : "Erreur inconnue");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="mb-4 flex flex-col items-center">
            <button
                className="px-4 py-2 bg-green-600 text-white rounded-lg shadow hover:bg-green-700 disabled:opacity-50"
                onClick={handleAutoApply}
                disabled={loading}
            >
                {loading ? "Application en cours..." : "Auto-appliquer tous les utilisateurs"}
            </button>
            {error && <div className="text-red-500 mt-2">{error}</div>}
            {success && <div className="text-green-600 mt-2">Tous les utilisateurs ont été appliqués !</div>}
        </div>
    );
};

// Match Result Modal Component
const MatchResultModal: React.FC<{
    match: any;
    onClose: () => void;
    onSubmit: (matchId: string, winnerId: string) => void;
    isOwner: boolean;
}> = ({ match, onClose, onSubmit, isOwner }) => {
    const [selectedWinner, setSelectedWinner] = useState<string>("");
    const [loading, setLoading] = useState(false);

    const handleSubmit = async () => {
        if (!selectedWinner) return;
        setLoading(true);
        await onSubmit(match.id, selectedWinner);
        setLoading(false);
        onClose();
    };

    const participants = Array.isArray(match.participants) ? match.participants : [];

    return (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
            <div className="bg-white rounded-xl p-6 max-w-md w-full mx-4">
                <h3 className="text-xl font-bold mb-4 text-blue-700">Résultat du match</h3>
                
                <div className="mb-4">
                    <p className="text-gray-600 mb-3">Qui a gagné ce match ?</p>
                    
                    {participants.map((participant: any) => {
                        const userObj = participant.user || participant;
                        return (
                            <label key={userObj.id} className="flex items-center mb-3 cursor-pointer">
                                <input
                                    type="radio"
                                    name="winner"
                                    value={userObj.id}
                                    checked={selectedWinner === userObj.id}
                                    onChange={(e) => setSelectedWinner(e.target.value)}
                                    className="mr-3"
                                    disabled={!isOwner}
                                />
                                <img
                                    src={userObj.profilePic || '/logo.png'}
                                    alt={userObj.name}
                                    className="w-8 h-8 rounded-full mr-3"
                                    onError={(e) => {
                                        const target = e.target as HTMLImageElement;
                                        if (target.src !== '/logo.png') {
                                            target.src = '/logo.png';
                                        }
                                    }}
                                />
                                <span className="font-medium">{userObj.name}</span>
                            </label>
                        );
                    })}
                </div>

                <div className="flex gap-3">
                    <button
                        onClick={onClose}
                        className="flex-1 px-4 py-2 bg-gray-300 text-gray-700 rounded-lg hover:bg-gray-400"
                    >
                        Annuler
                    </button>
                    {isOwner && (
                        <button
                            onClick={handleSubmit}
                            disabled={!selectedWinner || loading}
                            className="flex-1 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
                        >
                            {loading ? "..." : "Valider"}
                        </button>
                    )}
                </div>
            </div>
        </div>
    );
};

// Tournament Bracket Component with scoring system
const TournamentBracket: React.FC<{ 
    bracket: any[]; 
    onMatchResult: (matchId: string, winnerId: string) => void;
    isOwner: boolean;
}> = ({ bracket, onMatchResult, isOwner }) => {
    const [selectedMatch, setSelectedMatch] = useState<any | null>(null);
    
    const roundWidth = 300;
    const matchHeight = 140;
    const matchWidth = 240;
    
    // Calculate total dimensions
    const totalWidth = bracket.length * roundWidth;
    const maxMatchesInRound = Math.max(...bracket.map(round => round.matches.length));
    const totalHeight = maxMatchesInRound * matchHeight + 100;

    // Calculate positions for each match
    const getMatchPosition = (roundIdx: number, matchIdx: number, roundLength: number) => {
        const x = roundIdx * roundWidth + 30;
        const startY = (totalHeight - (roundLength * matchHeight)) / 2;
        const y = startY + matchIdx * matchHeight + 50;
        return { x, y };
    };

    // Check if a match can be played (both participants are present)
    const canPlayMatch = (match: any) => {
        if (!Array.isArray(match.participants)) return false;
        return match.participants.length === 2 && 
               match.participants.every((p: any) => (p.user || p).id !== 'TBD') &&
               !match.winner;
    };

    // Check if match is completed
    const isMatchCompleted = (match: any) => {
        return match.winner && match.status === 'completed';
    };

    return (
        <div className="relative" style={{ width: `${totalWidth}px`, height: `${totalHeight}px`, minHeight: '400px' }}>
            {/* SVG for connection lines */}
            <svg 
                className="absolute top-0 left-0 w-full h-full pointer-events-none" 
                style={{ zIndex: 1 }}
                width={totalWidth} 
                height={totalHeight}
            >
                {bracket.map((round, roundIdx) => {
                    if (roundIdx === 0) return null;
                    
                    return round.matches.map((match: any, matchIdx: number) => {
                        if (!match.sourceMatchIds || !Array.isArray(match.sourceMatchIds)) return null;
                        
                        const currentPos = getMatchPosition(roundIdx, matchIdx, round.matches.length);
                        
                        const sourcePositions = match.sourceMatchIds.map((sourceId: string) => {
                            const sourceMatch = bracket[roundIdx - 1].matches.find((m: any) => m.id === sourceId);
                            if (!sourceMatch) return null;
                            const sourceMatchIdx = bracket[roundIdx - 1].matches.indexOf(sourceMatch);
                            return getMatchPosition(roundIdx - 1, sourceMatchIdx, bracket[roundIdx - 1].matches.length);
                        }).filter(pos => pos !== null);

                        if (sourcePositions.length === 0) return null;

                        const connectorX = currentPos.x - 60;
                        const targetY = currentPos.y + 60;
                        
                        const sourceYs = sourcePositions.map(pos => pos.y + 60);
                        const minY = Math.min(...sourceYs);
                        const maxY = Math.max(...sourceYs);
                        const midY = (minY + maxY) / 2;

                        return (
                            <g key={`connection-group-${roundIdx}-${matchIdx}`}>
                                {sourcePositions.map((sourcePos, sourceIdx) => (
                                    <line
                                        key={`horizontal-${sourceIdx}`}
                                        x1={sourcePos.x + matchWidth}
                                        y1={sourcePos.y + 60}
                                        x2={connectorX}
                                        y2={sourcePos.y + 60}
                                        stroke="#3b82f6"
                                        strokeWidth="2"
                                        opacity="0.8"
                                    />
                                ))}
                                
                                {sourcePositions.length > 1 && (
                                    <line
                                        x1={connectorX}
                                        y1={minY}
                                        x2={connectorX}
                                        y2={maxY}
                                        stroke="#3b82f6"
                                        strokeWidth="2"
                                        opacity="0.8"
                                    />
                                )}
                                
                                <line
                                    x1={connectorX}
                                    y1={sourcePositions.length > 1 ? midY : sourceYs[0]}
                                    x2={connectorX + 30}
                                    y2={sourcePositions.length > 1 ? midY : sourceYs[0]}
                                    stroke="#3b82f6"
                                    strokeWidth="2"
                                    opacity="0.8"
                                />
                                
                                {((sourcePositions.length > 1 ? midY : sourceYs[0]) !== targetY) && (
                                    <line
                                        x1={connectorX + 30}
                                        y1={sourcePositions.length > 1 ? midY : sourceYs[0]}
                                        x2={connectorX + 30}
                                        y2={targetY}
                                        stroke="#3b82f6"
                                        strokeWidth="2"
                                        opacity="0.8"
                                    />
                                )}
                                
                                <line
                                    x1={connectorX + 30}
                                    y1={targetY}
                                    x2={currentPos.x}
                                    y2={targetY}
                                    stroke="#3b82f6"
                                    strokeWidth="2"
                                    opacity="0.8"
                                />
                            </g>
                        );
                    });
                })}
            </svg>

            {/* Match cards */}
            {bracket.map((round, roundIdx) => (
                <div key={`round-${roundIdx}`}>
                    {/* Round label */}
                    <div 
                        className="absolute font-bold text-blue-600 text-center"
                        style={{
                            left: `${roundIdx * roundWidth + 30}px`,
                            top: '10px',
                            width: `${matchWidth}px`
                        }}
                    >
                        {round.roundName || `Tour ${roundIdx + 1}`}
                    </div>
                    
                    {/* Matches */}
                    {round.matches.map((match: any, matchIdx: number) => {
                        const position = getMatchPosition(roundIdx, matchIdx, round.matches.length);
                        
                        let participants: any[] = [];
                        if (Array.isArray(match.participants)) {
                            participants = match.participants;
                        } else if (typeof match.name === 'string') {
                            // Parse team names from match.name (e.g., "Team1 vs Team2")
                            const teamNames = match.name.split(' vs ');
                            participants = teamNames.map((name: string) => ({ 
                                id: name.trim(), 
                                name: name.trim() 
                            }));
                        }

                        const completed = isMatchCompleted(match);
                        const playable = canPlayMatch(match);
                        const pending = !completed && !playable;

                        return (
                            <div
                                key={match.id || `match-${roundIdx}-${matchIdx}`}
                                onClick={() => {
                                    if (playable && isOwner) {
                                        setSelectedMatch(match);
                                    }
                                }}
                                className={`absolute rounded-xl shadow-lg border-2 transition-all duration-200 ${
                                    completed 
                                        ? 'bg-green-50 border-green-300 hover:border-green-400' 
                                        : playable && isOwner
                                            ? 'bg-yellow-50 border-yellow-300 hover:border-yellow-400 cursor-pointer'
                                            : pending
                                                ? 'bg-gray-50 border-gray-200'
                                                : 'bg-white border-blue-200 hover:border-blue-400'
                                }`}
                                style={{
                                    left: `${position.x}px`,
                                    top: `${position.y}px`,
                                    width: `${matchWidth}px`,
                                    height: '120px',
                                    zIndex: 2
                                }}
                            >
                                <div className="p-4 h-full flex flex-col justify-center relative">
                                    {/* Match status indicator */}
                                    <div className="absolute top-2 right-2">
                                        {completed && (
                                            <div className="w-3 h-3 bg-green-500 rounded-full" title="Match terminé"></div>
                                        )}
                                        {playable && !completed && (
                                            <div className="w-3 h-3 bg-yellow-500 rounded-full animate-pulse" title="Match en attente"></div>
                                        )}
                                        {pending && (
                                            <div className="w-3 h-3 bg-gray-400 rounded-full" title="En attente des qualifiés"></div>
                                        )}
                                    </div>

                                    {participants.length > 0 ? (
                                        <div className="space-y-3">
                                            {participants.map((participant: any, idx: number) => {
                                                const userObj = participant.user || participant;
                                                const isWinner = match.winner === userObj.id;
                                                const isTBD = userObj.id === 'TBD' || userObj.name === 'TBD';
                                                
                                                return (
                                                    <div 
                                                        key={userObj.id || userObj.name || idx} 
                                                        className={`flex items-center gap-3 p-2 rounded-lg transition-colors ${
                                                            isWinner ? 'bg-green-100 border border-green-300' : ''
                                                        }`}
                                                    >
                                                        <img
                                                            src={isTBD ? '/logo.png' : (userObj.profilePic || '/logo.png')}
                                                            alt={userObj.name || userObj.id || 'Player'}
                                                            className={`w-8 h-8 rounded-full border-2 shadow-sm flex-shrink-0 ${
                                                                isWinner ? 'border-green-500' : 'border-blue-300'
                                                            } ${isTBD ? 'opacity-50' : ''}`}
                                                            style={{ objectFit: 'cover' }}
                                                            onError={(e) => {
                                                                const target = e.target as HTMLImageElement;
                                                                if (target.src !== '/logo.png') {
                                                                    target.src = '/logo.png';
                                                                }
                                                            }}
                                                        />
                                                        <span className={`font-semibold text-sm truncate ${
                                                            isWinner ? 'text-green-800' : 'text-blue-900'
                                                        } ${isTBD ? 'text-gray-500' : ''}`}>
                                                            {isTBD ? 'En attente...' : (userObj.name || userObj.id || 'TBD')}
                                                            {isWinner && <span className="ml-2 text-green-600">👑</span>}
                                                        </span>
                                                    </div>
                                                );
                                            })}
                                            {participants.length === 2 && !completed && (
                                                <div className="text-center">
                                                    <span className="text-blue-500 font-bold text-xs">VS</span>
                                                </div>
                                            )}
                                        </div>
                                    ) : (
                                        <div className="text-center text-gray-500 text-sm">
                                            En attente...
                                        </div>
                                    )}

                                    {/* Action indicator for playable matches */}
                                    {playable && isOwner && (
                                        <div className="absolute bottom-2 left-1/2 transform -translate-x-1/2">
                                            <span className="text-xs text-yellow-600 font-medium">Cliquez pour noter</span>
                                        </div>
                                    )}
                                </div>
                            </div>
                        );
                    })}
                </div>
            ))}

            {/* Match Result Modal */}
            {selectedMatch && (
                <MatchResultModal
                    match={selectedMatch}
                    onClose={() => setSelectedMatch(null)}
                    onSubmit={onMatchResult}
                    isOwner={isOwner}
                />
            )}
        </div>
    );
};

const EventDetailPage = ({ params }: { params: Promise<{ id: string }> }) => {
    const { id } = React.use(params);
    const [event, setEvent] = useState<any>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const { user } = useUser();

    useEffect(() => {
        if (!id) return;
        const fetchEvent = async () => {
            try {
                setError(null);
                const res = await fetch(`/api/events/${id}`);
                if (!res.ok) throw new Error("Failed to fetch event");
                const data = await res.json();
                setEvent(data);
            } catch (err: any) {
                setError(err instanceof Error ? err.message : "Failed to fetch event");
            } finally {
                setLoading(false);
            }
        };
        fetchEvent();
    }, [id]);

    // Bracket state
    const [bracket, setBracket] = useState<any[][] | null>(null);
    const [bracketError, setBracketError] = useState<string | null>(null);
    const [bracketLoading, setBracketLoading] = useState(false);

    // Check if registration is closed
    const now = new Date();
    const applyEnd = event ? new Date(event.applyEnd) : null;
    const registrationClosed = event && applyEnd ? now > applyEnd : false;
    const isOwner = user && event && user.id === event.ownerId;

    useEffect(() => {
        if (event && registrationClosed && event.participants?.length > 1) {
            setBracketLoading(true);
            fetch(`/api/events/${event.id}/bracket`, { method: "POST" })
                .then(res => res.json())
                .then(data => {
                    if (data.error) setBracketError(data.error);
                    else setBracket(data.bracket);
                })
                .catch(() => setBracketError("Failed to fetch bracket"))
                .finally(() => setBracketLoading(false));
        } else {
            setBracket(null);
            setBracketError(null);
        }
    }, [event, registrationClosed]);

    // Handle match result submission
    const handleMatchResult = async (matchId: string, winnerId: string) => {
        if (!event || !isOwner) return;
        
        try {
            const res = await fetch(`/api/events/${event.id}/matches/${matchId}/result`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ winnerId })
            });
            
            if (!res.ok) throw new Error('Failed to update match result');
            
            // Refresh bracket data
            const bracketRes = await fetch(`/api/events/${event.id}/bracket`, { method: "POST" });
            const bracketData = await bracketRes.json();
            
            if (bracketData.error) {
                setBracketError(bracketData.error);
            } else {
                setBracket(bracketData.bracket);
                setBracketError(null);
            }
        } catch (error) {
            console.error('Error updating match result:', error);
            setBracketError('Failed to update match result');
        }
    };

    const handleAutoApplyRefresh = () => {
        if (id) {
            fetch(`/api/events/${id}`)
                .then(res => res.json())
                .then(data => setEvent(data))
                .catch(console.error);
        }
    };

    if (loading) return <div className="min-h-screen flex items-center justify-center">Loading...</div>;
    if (error) return <div className="min-h-screen flex items-center justify-center text-red-500">{error}</div>;
    if (!event) return <div className="min-h-screen flex items-center justify-center">Event not found</div>;

    return (
        <>
            <NavBar />
            <main className="min-h-screen bg-gradient-blue px-4 flex items-center justify-center pt-32">
                <div className="flex flex-col items-center w-full max-w-6xl">
                    {/* Auto-apply all users button (visible only for owner) */}
                    {isOwner && !registrationClosed && (
                        <AutoApplyAllUsers eventId={event.id} onApplied={handleAutoApplyRefresh} />
                    )}
                    
                    <header className="mb-10 text-center">
                        <h1 className="text-3xl md:text-5xl font-extrabold text-blue-700 mb-2 drop-shadow-sm">{event.name}</h1>
                        <p className="text-lg md:text-xl text-blue-500 font-medium">{event.description}</p>
                    </header>
                    
                    <section className="w-full bg-white/80 rounded-xl shadow-lg p-6 mb-8">
                        <div className="mb-4 text-blue-700 font-bold">Event ID: {event.id}</div>
                        <div className="mb-2 text-blue-600">Owner: {event.ownerId}</div>
                        <div className="mb-2 text-blue-600">Created: {new Date(event.createdAt).toLocaleString()}</div>
                        <div className="mb-2 text-blue-600">Registration: {new Date(event.applyStart).toLocaleString()} - {new Date(event.applyEnd).toLocaleString()}</div>
                        <div className="mb-2 text-blue-600">Participants: {event.participants?.length ?? 0}</div>
                    </section>
                    
                    {registrationClosed && (
                        <section className="w-full bg-white/90 rounded-xl shadow-lg p-6 mb-8">
                            <div className="flex justify-between items-center mb-4">
                                <h2 className="text-xl font-bold text-blue-700">Arbre du tournoi</h2>
                                {isOwner && (
                                    <div className="text-sm text-blue-600">
                                        <span className="inline-block w-3 h-3 bg-yellow-500 rounded-full mr-2"></span>
                                        Cliquez sur les matchs en attente pour noter le résultat
                                    </div>
                                )}
                            </div>
                            {bracketLoading ? (
                                <div className="text-blue-500">Génération de l'arbre...</div>
                            ) : bracketError ? (
                                <div className="text-red-500">{bracketError}</div>
                            ) : bracket ? (
                                <div className="overflow-auto">
                                    <TournamentBracket 
                                        bracket={bracket} 
                                        onMatchResult={handleMatchResult}
                                        isOwner={isOwner || false}
                                    />
                                </div>
                            ) : (
                                <div className="text-blue-400">Aucun arbre généré.</div>
                            )}
                        </section>
                    )}
                    
                    <button className="px-4 py-2 bg-blue-600 text-white rounded-lg shadow hover:bg-blue-700" onClick={() => window.history.back()}>
                        Back
                    </button>
                </div>
            </main>
        </>
    );
};

export default EventDetailPage;