<script lang="ts">
	import TailwindCss from './TailwindCSS.svelte'
	import logo from './assets/svelte.png'
	import Counter from './lib/Counter.svelte'
	import Nav from './lib/Nav.svelte'
	import Hero from './lib/Hero.svelte'
	import Team from './lib/Team.svelte'
	import Services from './lib/Services.svelte'
	import Footer from './lib/Footer.svelte'
	import Container from './lib/Container.svelte'
</script>

<TailwindCss />
<!-- class="max-w mx-auto px-4" -->
<main>
	<Container />
	<!-- <div class="pt-4 pb-12"> -->
	<!-- <Nav />
		<Hero />
		<Services />
		<Team />
		<Footer /> -->
	<!-- </div> -->
</main>

<style>
	:root {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
			Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
	}

	main {
		margin: 0 auto;
	}
</style>
