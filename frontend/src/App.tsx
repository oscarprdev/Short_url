import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ModalProvider } from './context/ModalProvider';
import Layout from './screens/Layout';

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
