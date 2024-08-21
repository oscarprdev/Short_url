import { mockedUrl } from '../../../tests/mocks/url-response';
import UrlCardCopyIcon from './UrlCardCopyIcon';
import { RenderResult, render } from '@testing-library/react';
import { afterEach, beforeEach, describe, it } from 'vitest';

describe('UrlCardCopyIcon', () => {
	let component: RenderResult;

	beforeEach(() => {
		component = render(<UrlCardCopyIcon url={mockedUrl.shortUrl} />);
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', () => {
		component.getByTestId('icon-copy');
	});
});
