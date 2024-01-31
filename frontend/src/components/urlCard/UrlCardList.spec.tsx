import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/react';
import UrlCardList from './UrlCardList';
import { mockedUrl } from '../../../tests/mocks/url-response';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

describe('UrlCardList', () => {
	it('Should render successfully', async () => {
		const component = render(
			<UrlCardList
				urls={[mockedUrl]}
				isHome={true}
				isRowsLayout={true}
			/>
		);
		component.getByRole('list');

		component.unmount();
	});

	it('Should have proper styles depending if isRowsLayout value is true', () => {
		const component = render(
			<UrlCardList
				urls={[mockedUrl]}
				isHome={true}
				isRowsLayout={true}
			/>
		);

		const list = component.getByRole('list');

		expect(list.className).toContain('w-[90vw] lg:w-[800px]');

		component.unmount();
	});

	it('Should have proper styles depending if isRowsLayout value is false', () => {
		const component = render(
			<UrlCardList
				urls={[mockedUrl]}
				isHome={true}
				isRowsLayout={false}
			/>
		);

		const list = component.getByRole('list');

		expect(list.className).toContain('w-[90vw] lg:w-[1100px]');

		component.unmount();
	});

	it('Should display normal url card if isHome is true', () => {
		const component = render(
			<UrlCardList
				urls={[mockedUrl]}
				isHome={true}
				isRowsLayout={false}
			/>
		);

		component.getByRole('listitem');

		component.getByTestId(`url-link-${mockedUrl.id}`);

		component.unmount();
	});

	it('Should display normal url card if isHome is false', () => {
		const queryClient = new QueryClient();
		const component = render(
			<QueryClientProvider client={queryClient}>
				<UrlCardList
					urls={[mockedUrl]}
					isHome={false}
					isRowsLayout={false}
				/>
			</QueryClientProvider>
		);

		component.getByRole('listitem');

		component.getByDisplayValue(mockedUrl.titleUrl);

		component.unmount();
	});
});
