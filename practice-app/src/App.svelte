<script>
  import { shapes, generateRandomNumber as generateShot } from './utils.js';

  let randomYardage = null;
  let randomShape = null;
  let randomObstruction = null;
  let mode = 'yardage';

  function generateRandomNumber() {
    const result = generateShot(mode, shapes);
    randomYardage = result.randomYardage;
    randomShape = result.randomShape;
    randomObstruction = result.randomObstruction;
  }
  
  function setMode(newMode) {
    mode = newMode;
    randomYardage = null;
    randomShape = null;
    randomObstruction = null;
  }
</script>

<main>
  <h1>Practice App</h1>
  
  <div class="mode-selector">
    <button 
      class="mode-button"
      class:active={mode === 'yardage'}
      on:click={() => setMode('yardage')}
    >
      Yardage
    </button>
    <button 
      class="mode-button"
      class:active={mode === 'yardage-shape'}
      on:click={() => setMode('yardage-shape')}
    >
      Yardage + Shape
    </button>
    <button 
      class="mode-button"
      class:active={mode === 'course'}
      on:click={() => setMode('course')}
    >
      Course
    </button>
  </div>

  {#if mode === 'yardage'}
    {#if randomYardage !== null}
      <p class="random-number">Yardage: <strong>{randomYardage}</strong></p>
    {:else}
      <p>Click the button to generate a random yardage</p>
    {/if}
  {:else if mode === 'yardage-shape'}
    {#if randomYardage !== null && randomShape !== null}
      <div class="random-display">
        <p class="random-number">Yardage: <strong>{randomYardage}</strong></p>
        <p class="random-shape">Shape: <strong>{randomShape}</strong></p>
      </div>
    {:else}
      <p>Click the button to generate a random yardage and shape</p>
    {/if}
  {:else if mode === 'course'}
    {#if randomYardage !== null && randomShape !== null && randomObstruction !== null}
      <div class="random-display">
        <p class="random-number">Yardage: <strong>{randomYardage}</strong></p>
        <p class="random-shape">Shape: <strong>{randomShape}</strong></p>
        <p class="random-obstruction">Obstruction: <strong>{randomObstruction}</strong></p>
      </div>
    {:else}
      <p>Click the button to generate a random yardage, shape, and obstruction</p>
    {/if}
  {/if}
  
  <button on:click={generateRandomNumber}>Generate Shot</button>
</main>
