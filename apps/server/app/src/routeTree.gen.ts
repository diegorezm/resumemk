/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file was automatically generated by TanStack Router.
// You should NOT make any changes in this file as it will be overwritten.
// Additionally, you should also exclude this file from your linter and/or formatter to prevent it from being checked or modified.

import { createFileRoute } from '@tanstack/react-router'

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as DocumentsIndexImport } from './routes/documents/index'
import { Route as DocumentsDocumentIdImport } from './routes/documents/$documentId'

// Create Virtual Routes

const IndexLazyImport = createFileRoute('/')()

// Create/Update Routes

const IndexLazyRoute = IndexLazyImport.update({
  id: '/',
  path: '/',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/index.lazy').then((d) => d.Route))

const DocumentsIndexRoute = DocumentsIndexImport.update({
  id: '/documents/',
  path: '/documents/',
  getParentRoute: () => rootRoute,
} as any)

const DocumentsDocumentIdRoute = DocumentsDocumentIdImport.update({
  id: '/documents/$documentId',
  path: '/documents/$documentId',
  getParentRoute: () => rootRoute,
} as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      id: '/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof IndexLazyImport
      parentRoute: typeof rootRoute
    }
    '/documents/$documentId': {
      id: '/documents/$documentId'
      path: '/documents/$documentId'
      fullPath: '/documents/$documentId'
      preLoaderRoute: typeof DocumentsDocumentIdImport
      parentRoute: typeof rootRoute
    }
    '/documents/': {
      id: '/documents/'
      path: '/documents'
      fullPath: '/documents'
      preLoaderRoute: typeof DocumentsIndexImport
      parentRoute: typeof rootRoute
    }
  }
}

// Create and export the route tree

export interface FileRoutesByFullPath {
  '/': typeof IndexLazyRoute
  '/documents/$documentId': typeof DocumentsDocumentIdRoute
  '/documents': typeof DocumentsIndexRoute
}

export interface FileRoutesByTo {
  '/': typeof IndexLazyRoute
  '/documents/$documentId': typeof DocumentsDocumentIdRoute
  '/documents': typeof DocumentsIndexRoute
}

export interface FileRoutesById {
  __root__: typeof rootRoute
  '/': typeof IndexLazyRoute
  '/documents/$documentId': typeof DocumentsDocumentIdRoute
  '/documents/': typeof DocumentsIndexRoute
}

export interface FileRouteTypes {
  fileRoutesByFullPath: FileRoutesByFullPath
  fullPaths: '/' | '/documents/$documentId' | '/documents'
  fileRoutesByTo: FileRoutesByTo
  to: '/' | '/documents/$documentId' | '/documents'
  id: '__root__' | '/' | '/documents/$documentId' | '/documents/'
  fileRoutesById: FileRoutesById
}

export interface RootRouteChildren {
  IndexLazyRoute: typeof IndexLazyRoute
  DocumentsDocumentIdRoute: typeof DocumentsDocumentIdRoute
  DocumentsIndexRoute: typeof DocumentsIndexRoute
}

const rootRouteChildren: RootRouteChildren = {
  IndexLazyRoute: IndexLazyRoute,
  DocumentsDocumentIdRoute: DocumentsDocumentIdRoute,
  DocumentsIndexRoute: DocumentsIndexRoute,
}

export const routeTree = rootRoute
  ._addFileChildren(rootRouteChildren)
  ._addFileTypes<FileRouteTypes>()

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/",
        "/documents/$documentId",
        "/documents/"
      ]
    },
    "/": {
      "filePath": "index.lazy.tsx"
    },
    "/documents/$documentId": {
      "filePath": "documents/$documentId.tsx"
    },
    "/documents/": {
      "filePath": "documents/index.tsx"
    }
  }
}
ROUTE_MANIFEST_END */
