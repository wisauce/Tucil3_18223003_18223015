document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('solveForm');
    const algorithmSelect = document.getElementById('algorithm');
    const heuristicGroup = document.getElementById('heuristicGroup');
    const solveBtn = document.getElementById('solveBtn');
    const btnText = document.getElementById('btnText');
    const resultPanel = document.getElementById('resultPanel');
    const downloadBtn = document.getElementById('downloadBtn');
    
    // Playback elements
    const emptyStateMsg = document.getElementById('emptyStateMsg');
    const boardContainer = document.getElementById('boardContainer');
    const boardGrid = document.getElementById('boardGrid');
    const actor = document.getElementById('actor');
    const playbackControls = document.getElementById('playbackControls');
    const stepSlider = document.getElementById('stepSlider');
    const stepCounter = document.getElementById('stepCounter');
    const btnPlay = document.getElementById('btnPlay');
    const iconPlay = document.getElementById('iconPlay');
    const iconPause = document.getElementById('iconPause');
    const btnPrev = document.getElementById('btnPrev');
    const btnNext = document.getElementById('btnNext');

    let currentResponse = null;
    let currentStep = 0;
    let isPlaying = false;
    let playInterval = null;
    let boardWidth = 0;
    const TILE_SIZE = 40; // 40px width/height
    const GAP_SIZE = 4;   // gap-1 in tailwind = 4px

    // Toggle heuristic dropdown
    algorithmSelect.addEventListener('change', (e) => {
        const val = e.target.value;
        if (val === 'gbfs' || val === 'a*') {
            heuristicGroup.classList.remove('hidden');
        } else {
            heuristicGroup.classList.add('hidden');
        }
    });

    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        
        const fileInput = document.getElementById('fileInput');
        if (!fileInput.files.length) return;

        const formData = new FormData();
        formData.append('file', fileInput.files[0]);
        formData.append('algorithm', algorithmSelect.value);
        formData.append('heuristic', document.getElementById('heuristic').value);

        // Reset UI
        btnText.textContent = '⏳ Solving...';
        solveBtn.disabled = true;
        resultPanel.classList.add('hidden');
        playbackControls.classList.add('hidden');
        emptyStateMsg.classList.remove('hidden');
        boardContainer.classList.add('hidden');
        stepCounter.classList.add('hidden');
        pausePlayback();

        try {
            const res = await fetch('/api/solve', {
                method: 'POST',
                body: formData
            });
            const data = await res.json();
            
            if (!data.success) {
                alert(data.message);
                return;
            }

            currentResponse = data;
            
            // Fill result panel
            document.getElementById('resPath').textContent = data.path;
            document.getElementById('resCost').textContent = data.states[data.states.length-1].Cost;
            document.getElementById('resTime').textContent = data.time_ms + ' ms';
            document.getElementById('resIter').textContent = data.iterations;
            resultPanel.classList.remove('hidden');

            // Setup Board
            setupBoard(data);
            
            // Setup Playback
            currentStep = 0;
            stepSlider.max = data.states.length - 1;
            stepSlider.value = 0;
            playbackControls.classList.remove('hidden');
            emptyStateMsg.classList.add('hidden');
            boardContainer.classList.remove('hidden');
            stepCounter.classList.remove('hidden');
            
            updateBoardState(0);

        } catch (err) {
            alert('Terjadi kesalahan koneksi ke server.');
            console.error(err);
        } finally {
            btnText.textContent = 'Solve Puzzle';
            solveBtn.disabled = false;
        }
    });

    function setupBoard(data) {
        boardWidth = data.board_width;
        const boardHeight = data.board_height;
        
        boardGrid.style.gridTemplateColumns = `repeat(${boardWidth}, ${TILE_SIZE}px)`;
        boardGrid.style.gridTemplateRows = `repeat(${boardHeight}, ${TILE_SIZE}px)`;
        boardGrid.innerHTML = '';

        for (let y = 0; y < boardHeight; y++) {
            for (let x = 0; x < boardWidth; x++) {
                const char = data.board[y][x];
                const tile = document.createElement('div');
                tile.className = 'board-tile text-sm ';
                
                if (char === 'X') {
                    tile.classList.add('bg-slate-800', 'text-slate-800'); // Solid wall
                } else if (char === 'L') {
                    tile.classList.add('bg-orange-500', 'text-white');
                    tile.textContent = 'L';
                } else if (char === '*') {
                    tile.classList.add('bg-slate-200');
                } else if (char >= '0' && char <= '9') {
                    tile.classList.add('bg-emerald-500', 'text-white');
                    tile.textContent = char;
                    // Save id to easily dim it later
                    tile.id = `tile-${x}-${y}`;
                } else if (char === 'O') {
                    tile.classList.add('bg-emerald-700', 'text-white', 'rounded-full');
                    tile.textContent = 'O';
                } else if (char === 'Z') {
                    tile.classList.add('bg-slate-200'); // Initial Z spot becomes empty
                } else {
                    tile.classList.add('bg-slate-200');
                }
                
                boardGrid.appendChild(tile);
            }
        }
    }

    function updateBoardState(stepIdx) {
        if (!currentResponse) return;
        const state = currentResponse.states[stepIdx];
        
        // Update actor position
        const posX = state.X * (TILE_SIZE + GAP_SIZE);
        const posY = state.Y * (TILE_SIZE + GAP_SIZE);
        actor.style.transform = `translate(${posX}px, ${posY}px)`;

        // Reset all tiles to default, then dim collected numbers
        const boardHeight = currentResponse.board_height;
        for (let y = 0; y < boardHeight; y++) {
            for (let x = 0; x < boardWidth; x++) {
                const char = currentResponse.board[y][x];
                if (char >= '0' && char <= '9') {
                    const t = document.getElementById(`tile-${x}-${y}`);
                    if (t) {
                        if (parseInt(char) < state.NextNumber) {
                            t.classList.replace('bg-emerald-500', 'bg-slate-300');
                            t.classList.replace('text-white', 'text-slate-400');
                        } else {
                            t.classList.replace('bg-slate-300', 'bg-emerald-500');
                            t.classList.replace('text-slate-400', 'text-white');
                        }
                    }
                }
            }
        }

        stepSlider.value = stepIdx;
        stepCounter.textContent = `Step ${stepIdx} / ${currentResponse.states.length - 1}`;
    }

    function playNext() {
        if (currentStep < currentResponse.states.length - 1) {
            currentStep++;
            updateBoardState(currentStep);
        } else {
            pausePlayback();
        }
    }

    function playPrev() {
        if (currentStep > 0) {
            currentStep--;
            updateBoardState(currentStep);
        }
    }

    function togglePlayback() {
        if (isPlaying) {
            pausePlayback();
        } else {
            if (currentStep >= currentResponse.states.length - 1) {
                currentStep = 0; 
                updateBoardState(0);
            }
            isPlaying = true;
            if (iconPlay && iconPause) {
                iconPlay.classList.add('hidden');
                iconPause.classList.remove('hidden');
            }
            // Wait a tiny bit if resetting from end before sliding
            setTimeout(() => {
                if(!isPlaying) return;
                playInterval = setInterval(playNext, 400); 
            }, 50);
        }
    }

    function pausePlayback() {
        isPlaying = false;
        if (iconPlay && iconPause) {
            iconPlay.classList.remove('hidden');
            iconPause.classList.add('hidden');
        }
        if (playInterval) clearInterval(playInterval);
    }

    // Event Listeners for Playback
    btnPlay.addEventListener('click', togglePlayback);
    btnNext.addEventListener('click', () => { pausePlayback(); playNext(); });
    btnPrev.addEventListener('click', () => { pausePlayback(); playPrev(); });
    stepSlider.addEventListener('input', (e) => {
        pausePlayback();
        currentStep = parseInt(e.target.value);
        updateBoardState(currentStep);
    });

    // Download functionality
    downloadBtn.addEventListener('click', () => {
        if (!currentResponse) return;
        
        let content = `[SOLUSI ICE SLIDING PUZZLE]\n`;
        content += `Algoritma: ${algorithmSelect.options[algorithmSelect.selectedIndex].text}\n`;
        if(algorithmSelect.value === 'gbfs' || algorithmSelect.value === 'a*') {
            const h = document.getElementById('heuristic');
            content += `Heuristik: ${h.options[h.selectedIndex].text}\n`;
        }
        content += `Rute: ${currentResponse.path}\n`;
        content += `Cost Total: ${currentResponse.states[currentResponse.states.length-1].Cost}\n`;
        content += `Waktu Eksekusi: ${currentResponse.time_ms} ms\n`;
        content += `Banyak Iterasi: ${currentResponse.iterations}\n\n`;
        
        content += `[LANGKAH-LANGKAH]\n`;
        for (let i = 0; i < currentResponse.states.length; i++) {
            const state = currentResponse.states[i];
            const move = i === 0 ? "START" : currentResponse.path[i-1];
            content += `Step ${i} : ${move}\n`;
            
            // Draw board string
            for (let y = 0; y < currentResponse.board_height; y++) {
                let rowStr = "";
                for (let x = 0; x < currentResponse.board_width; x++) {
                    if (x === state.X && y === state.Y) {
                        rowStr += "Z";
                        continue;
                    }
                    const char = currentResponse.board[y][x];
                    if (char >= '0' && char <= '9' && parseInt(char) < state.NextNumber) {
                        rowStr += "*";
                    } else if (char === 'Z') {
                        rowStr += "*"; 
                    } else {
                        rowStr += char;
                    }
                }
                content += rowStr + "\n";
            }
            content += "\n";
        }

        const blob = new Blob([content], { type: "text/plain;charset=utf-8" });
        const url = URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        const fn = document.getElementById('fileInput').files[0].name.split('.')[0];
        a.download = `solusi_${fn}_${algorithmSelect.value}.txt`;
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);
        URL.revokeObjectURL(url);
    });
});
