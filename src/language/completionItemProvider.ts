import {
  getStandardSQLCompletionProvider,
  LanguageCompletionProvider,
} from '@grafana/experimental';
import { MACROS } from './macros';

interface CompletionProviderGetterArgs {
  getEventDataStores: () => Promise<EventDataStoreDefinition[]>;
}

export interface EventDataStoreDefinition {
  name: string;
  completion?: string;
}

export const getCtlCompletionProvider: (args: CompletionProviderGetterArgs) => LanguageCompletionProvider =
  ({ getEventDataStores }) =>
  (monaco, language) => {
    return {
      // get standard SQL completion provider which will resolve functions and macros
      ...(language && getStandardSQLCompletionProvider(monaco, language)),
      triggerCharacters: ['.', ' ', '$', ',', '(', "'"],
      eventDataStores: {
        resolve: getEventDataStores,
      },
      supportedMacros: () => MACROS,
    };
  };
