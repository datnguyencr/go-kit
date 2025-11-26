async function fetchPokemon() {
    try {
        const res = await fetch('/api/pokemon');
        const data = await res.json();
        const container = document.getElementById('pokemon-list');

        // Build all cards as a single string
        const cardsHTML = data.map(p => `
            <div class="pokemon-card">
                <div class="pokemon-name">${p.name}</div>
                <div class="pokedex-number">#${p.pokedex_number}</div>
                <div class="type-container">
                    <span class="type-${p.type1}">${p.type1.toUpperCase()}</span>
                    ${p.type2 ? `<span class="type-${p.type2}">${p.type2.toUpperCase()}</span>` : ''}
                </div>
                <p class="classification">Classification: ${p.classification}</p>

                <h3>Base Stats</h3>
                <div class="stat-line"><span>HP:</span> <span class="stat-value">${p.hp}</span></div>
                <div class="stat-line"><span>Attack:</span> <span class="stat-value">${p.attack}</span></div>
                <div class="stat-line"><span>Defense:</span> <span class="stat-value">${p.defense}</span></div>
                <div class="stat-line"><span>Sp. Atk:</span> <span class="stat-value">${p.sp_attack}</span></div>
                <div class="stat-line"><span>Sp. Def:</span> <span class="stat-value">${p.sp_defense}</span></div>
                <div class="stat-line"><span>Speed:</span> <span class="stat-value">${p.speed}</span></div>
                <div class="stat-line"><span>Base Total:</span> <span class="stat-value">${p.base_total}</span></div>

                <h3>Other Info</h3>
                <div class="stat-line"><span>Abilities:</span> <span class="stat-value">${p.abilities.join(', ')}</span></div>
                <div class="stat-line"><span>Height:</span> <span class="stat-value">${p.height} m</span></div>
                <div class="stat-line"><span>Weight:</span> <span class="stat-value">${p.weight} kg</span></div>
            </div>
        `).join('');

        container.innerHTML = cardsHTML;
    } catch (err) {
        console.error('Error fetching Pokémon data:', err);
    }
}


document.addEventListener('DOMContentLoaded', fetchPokemon);
