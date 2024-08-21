import { ModalProvider } from './context/ModalProvider';
import Layout from './screens/Layout';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

function App() {
	const queryClient = new QueryClient();

	return (
		<QueryClientProvider client={queryClient}>
			<ModalProvider>
				<Layout />
			</ModalProvider>
		</QueryClientProvider>
	);
}

export default App;
