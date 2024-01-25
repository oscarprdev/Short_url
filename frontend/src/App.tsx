import Footer from './components/Footer';
import Home from './screens/Home';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { Route } from 'wouter';
import RedirectComponent from './screens/Redirect';
import { ModalProvider } from './context/ModalProvider';

function App() {
	const queryClient = new QueryClient();

	return (
		<QueryClientProvider client={queryClient}>
			<ModalProvider>
				<main className='relative flex flex-col items-center h-screen w-screen overflow-hidden bg-black'>
					<Route
						path='/'
						component={Home}
					/>
					<Route
						path='/:id'
						component={RedirectComponent}
					/>
					<div className='absolute bottom-0 left-0 right-0 top-0 bg-[linear-gradient(to_right,#4f4f4f2e_1px,transparent_1px),linear-gradient(to_bottom,#8080800a_1px,transparent_1px)] bg-[size:14px_24px]'></div>
					<div className='absolute left-0 right-0 top-[-10%] h-[1000px] w-[1000px] rounded-full bg-[radial-gradient(circle_400px_at_50%_300px,#fbfbfb36,#00000080)]'></div>
					<Footer />
				</main>
			</ModalProvider>
		</QueryClientProvider>
	);
}

export default App;
